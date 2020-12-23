from concurrent import futures

import grpc
import validators

from . import twitter_pb2 as pb2
from . import twitter_pb2_grpc as pb2_grpc
from orakarik.scrape.tweet_scrape import SnTweetScrape
from .utils import date as date_parser


class TwitterStream(pb2_grpc.twitterServicer):

    def StreamRequest(self, request, context):
        result = {'message': 'ok', 'updateAt': '123', 'items': []}
        keyword = validators.length(request.keyword, min=1, max=120)
        search_type = validators.length(request.search_type, min=1, max=25)
        since = validators.length(request.since, min=1, max=30)
        until = validators.length(request.until, min=1, max=30)

        if not keyword:
            result['message'] = f'keyword length must min {keyword.min} max {keyword.max}'
            return pb2.twitterResponse(**result)
        if not search_type:
            result['message'] = f'search_type length must min {search_type.min} max {search_type.max}'
            return pb2.twitterResponse(**result)
        if not since:
            result['message'] = f'since length must min {since.min} max {since.max}'
            return pb2.twitterResponse(**result)
        if not until:
            result['message'] = f'until length must min {until.min} max {until.max}'
            return pb2.twitterResponse(**result)

        if not keyword and not search_type and not since and not until:
            result['message'] = f'[keyword, search_type, since, until] required'
            return pb2.twitterResponse(**result)

        searchType = str(request.search_type)

        filtered = date_parser.DateSettings(request.since, request.until)
        filteredO = filtered.get_date()
        result['updateAt'] = filteredO['step']

        with_proxy = True
        if not request.proxy or request.proxy == "":
            with_proxy = False

        proxy_host = "proxy.crawlera.com"
        proxy_port = "8010"
        proxy_auth = f"{request.proxy}:"
        proxies = {
            "http": "http://{}@{}:{}/".format(proxy_auth, proxy_host, proxy_port),
            "https": "https://{}@{}:{}/".format(proxy_auth, proxy_host, proxy_port),
        }

        try:
            scrape = SnTweetScrape(filteredO['since'], filteredO['until'], filteredO['count'],
                                   proxy=with_proxy,
                                   proxy_dict=proxies)
            if searchType == "account":
                items = scrape.tweetAccount(request.keyword.replace("@", ""), lang="id")
            elif searchType == "mention":
                text = request.keyword
                if "@" not in text:
                    text = "@" + text

                items = scrape.tweetSearch(text, lang="id")
            elif searchType == "hashtag":
                text = request.keyword
                if "#" in request.keyword:
                    text = request.keyword.replace("#", "")
                items = scrape.tweetHashtag(text, lang="id")
            else:
                items = scrape.tweetSearch(request.keyword, lang="id")

            return pb2.twitterResponse(message=result['message'], updateAt=result['updateAt'], items=items)
        except TypeError as e:
            print(e)
            return pb2.twitterResponse(message='error', updateAt='', items=[])


def serve(ports):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=8))
    pb2_grpc.add_twitterServicer_to_server(TwitterStream(), server)
    port = server.add_insecure_port(f'0.0.0.0:{ports}')
    print("Twitter port at {}".format(port))
    server.start()
    server.wait_for_termination()
