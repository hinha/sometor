from __future__ import print_function
from decouple import config

import datetime
import os
import re
import string
import time

import redis

from orakarik.scrape.cleansing import cleansing
from orakarik.scrape.sentistrength import sentistrength
from . import twitter


# Create an Api instance.
api = twitter.Api(consumer_key=config("TWITTER_CONSUMER_KEY"),
                  consumer_secret=config("TWITTER_CONSUMER_SECRET"),
                  access_token_key=config("TWITTER_ACCESS_TOKEN"),
                  access_token_secret=config("TWITTER_ACCESS_TOKEN_SECRET"))

senti = sentistrength.sentistrength()

client = redis.StrictRedis(host=config("REDIS_HOST"), port=6379, db=2)


# client = redis.StrictRedis("127.0.0.1",port=6379, db=10)
class TweetSearch(object):

    def searchMention(self, term, since, until, count=5):
        feeds = []
        results = api.GetSearch(term=term, since=since, until=until, count=count, locale="in", include_entities=True)
        statusesV = {}
        for result in results:
            try:
                statusesV = self.__Mentiongettiemline_cache(result.id)
            except:
                dateSec = time.mktime(datetime.datetime.strptime(result.created_at, "%a %b %d %X +0000 %Y").timetuple())
                formatted_date = datetime.datetime.fromtimestamp(dateSec).strftime("%Y-%m-%d %H:%M")
                created_at_time = str(datetime.datetime.fromtimestamp(dateSec))[11:16]
                text_norm = self.__remove_url(result.full_text)
                # replace puctuantion
                text_norm = self.__remove_punctuation(text_norm)
                text_clean = cleansing.clean_text_light(text_norm)
                sentiment = senti.main(text_clean)
                statusesV = {
                    "screen_name": result.user.name,
                    "user_verified": result.user.verified,
                    "user_name": result.user.screen_name,
                    "full_text_norm": self.__remove_url(result.full_text),
                    "full_text_clean": text_clean,
                    "text_sentiment": sentiment,
                    "created_at_time": created_at_time,
                    "timestamp": int(dateSec),
                    "created_at": formatted_date,
                    "favorite_count": result.favorite_count,
                    "full_text": result.full_text,
                    "hashtags": ["#" + x.text for x in result.hashtags if x],
                    "mentions": ["@" + x.screen_name for x in result.user_mentions if x],
                    "id": result.id,
                    "lang": result.lang,
                    "retweet_count": result.retweet_count,
                    "permalink": f"https://twitter.com/{result.user.id}/status/{result.id}",
                    "engagement": result.retweet_count + result.favorite_count + result.user.statuses_count,

                    "user_description": result.user.description,
                    "user_location": result.user.location,
                    "user_favourites_count": result.user.favourites_count,
                    "user_followers_count": result.user.followers_count,
                    "user_friends_count": result.user.friends_count,
                    "user_geo_enabled": result.user.geo_enabled,
                    "user_id": result.user.id,
                    "user_listed_count": result.user.listed_count,
                    "user_statuses_count": result.user.statuses_count,
                    "user_url": result.user.url if result.user.url else "",
                    "profile_image_url_https": result.user.profile_image_url_https,
                }

                self.__Mentionsettimeline_cache(result.id, statusesV)
            feeds.append(statusesV)

        return feeds

    def searchAccounts(self, name, limit=5):
        page = 0
        count = 0
        # users = api.GetUsersSearch(term=name, page=1, count=10)
        users = api.GetUser(screen_name=name)
        views = {
            "user_verified": users.verified,
            "screen_name": users.name,
            "user_name": users.screen_name,
            "user_description": users.description,
            "user_location": users.location,
            "user_favourites_count": users.favourites_count,
            "user_followers_count": users.followers_count,
            "user_friends_count": users.friends_count,
            "user_geo_enabled": users.geo_enabled,
            "user_id": users.id,
            "user_listed_count": users.listed_count,
            "user_statuses_count": users.statuses_count,
            "user_url": users.url if users.url else "",
            "user_created_at": users.created_at,
            "user_profile_image_url": users.profile_image_url_https,
        }
        users_timeline = api.GetUserTimeline(user_id=users.id, exclude_replies=True, count=limit + 1)

        feeds = []
        for tt in users_timeline:
            fodateSec = time.mktime(datetime.datetime.strptime(tt.created_at, "%a %b %d %X +0000 %Y").timetuple())
            # print(self.tanggal_start, fodateSec, self.tanggal_end)
            statusesV = {}
            count = count + 1

            try:
                statusesV = self.__gettiemline_cache(tt.id)
            except Exception as e:
                count += 1
                statuses = api.GetStatus(status_id=tt.id)

                dateSec = time.mktime(
                    datetime.datetime.strptime(statuses.created_at, "%a %b %d %X +0000 %Y").timetuple())
                formatted_date = datetime.datetime.fromtimestamp(dateSec).strftime("%Y-%m-%d %H:%M")
                created_at_time = str(datetime.datetime.fromtimestamp(dateSec))[11:16]
                text_norm = self.__remove_url(statuses.full_text)
                # replace puctuantion
                text_norm = self.__remove_punctuation(text_norm)
                text_clean = cleansing.clean_text_light(text_norm)
                sentiment = senti.main(text_clean)
                statusesV = {
                    "full_text_norm": self.__remove_url(statuses.full_text),
                    "full_text_clean": text_clean,
                    "text_sentiment": sentiment,
                    "created_at_time": created_at_time,
                    "timestamp": int(dateSec),
                    "created_at": formatted_date,
                    "favorite_count": statuses.favorite_count,
                    "full_text": statuses.full_text,
                    "hashtags": ["#" + x.text for x in statuses.hashtags if x],
                    "mentions": ["@" + x.screen_name for x in statuses.user_mentions if x],
                    "id": statuses.id,
                    "lang": statuses.lang,
                    "retweet_count": statuses.retweet_count,
                    "permalink": f"https://twitter.com/{statuses.user.id}/status/{statuses.id}",
                    "engagement": statuses.retweet_count + statuses.favorite_count + statuses.user.statuses_count
                }
                statusesV.update(views)

                self.__settimeline_cache(tt.id, statusesV)

            feeds.append(statusesV)

            if page % 25 == 0:
                time.sleep(1)

            if count >= limit != 0:
                break

        return feeds

    def __Mentiongettiemline_cache(self, key) -> dict:
        data = client.get(f"mention:{key}")
        results = data.decode("utf-8")

        return eval(results)

    def __Mentionsettimeline_cache(self, key, raw: dict, expired=86400) -> None:
        client.set(f"mention:{key}", str(raw), ex=expired)

    def __gettiemline_cache(self, key) -> dict:
        pprint = f"scrape:twitter_statuses:{key}"
        data = client.get(pprint)
        results = data.decode("utf-8")

        return eval(results)

    def __settimeline_cache(self, key, raw: dict, expired=86400) -> None:
        pprint = f"scrape:twitter_statuses:{key}"
        client.set(pprint, str(raw), ex=expired)

    def __remove_punctuation(self, s):  # punctuation
        for c in string.punctuation:
            s = s.replace(c, " ")

        # replace phone number
        comp_number = re.compile(
            "\d?(\(?\d{3}\D{0,3}\d{3}\D{0,3}\d{4})", re.S)
        text = comp_number.sub(
            lambda m: re.sub('\d', ' ', m.group(1)), s
        )
        text = self.__concate_duplicate(text.lower())
        return text

    def __remove_url(self, text):
        # remove all url
        text = re.sub(r" ?(f|ht)(tp)(s?)(://)(.*)[.|/](.*)", "", text)

        # remove email
        text = re.sub(r"[\w]+@[\w]+\.[c][o][m]", "", text)
        # remove text twit
        text = re.sub(r'((pic\.[^\s]+)|(twitter))', '', text)
        # remove mentions, hashtag and web
        text = re.sub(r"(?:\@|#|http?\://)\S+", "", text)
        # remove url
        text = re.sub(
            r'((www\.[^\s]+)|(https?://[^\s]+))', '', text)
        text = re.sub(r'((https?://[^\s]+))', '', text)
        text = re.sub(r"(pic[^\s]+)|[\w]+\.[c][o][m]", "", text)
        # replace non ascii
        text = re.sub(r'[^\x00-\x7F]+', ' ', text)

        # Remove additional white spaces
        text = re.sub('[\s]+', ' ', text)
        text = re.sub('[\n]+', ' ', text)

        return text

    def __remove_emojis(self, data):
        emoj = re.compile("["
                          u"\U0001F600-\U0001F64F"  # emoticons
                          u"\U0001F300-\U0001F5FF"  # symbols & pictographs
                          u"\U0001F680-\U0001F6FF"  # transport & map symbols
                          u"\U0001F1E0-\U0001F1FF"  # flags (iOS)
                          u"\U00002500-\U00002BEF"  # chinese char
                          u"\U00002702-\U000027B0"
                          u"\U00002702-\U000027B0"
                          u"\U000024C2-\U0001F251"
                          u"\U0001f926-\U0001f937"
                          u"\U00010000-\U0010ffff"
                          u"\u2640-\u2642"
                          u"\u2600-\u2B55"
                          u"\u200d"
                          u"\u23cf"
                          u"\u23e9"
                          u"\u231a"
                          u"\ufe0f"  # dingbats
                          u"\u3030"
                          "]+", re.UNICODE)
        return re.sub(emoj, '', data)

    def __concate_duplicate(self, text):
        te = "a" + r"{3}"
        rep = re.sub(te, " 3", text)
        te = "i" + r"{3}"
        rep = re.sub(te, " 3", rep)
        te = "u" + r"{3}"
        rep = re.sub(te, " 3", rep)
        te = "e" + r"{3}"
        rep = re.sub(te, " 3", rep)
        te = "o" + r"{3}"
        rep = re.sub(te, " 3", rep)

        # Extends
        te = "c" + r"{3}"
        rep = re.sub(te, " 3", rep)
        te = "k" + r"{3}"
        rep = re.sub(te, " 3", rep)
        te = "w" + r"{3}"
        rep = re.sub(te, " 3", rep)
        te = "h" + r"{3}"
        rep = re.sub(te, " 3", rep)

        return rep


def check_user(name):
    try:
        users = api.GetUser(screen_name=name)
        if users:
            if users.protected is True:
                return {'status': False, 'is_private': 1}
            return {'status': True, 'is_private': 0, "id": users.id}
        else:
            return {'status': False, 'is_private': 0, "id": "00"}
    except Exception as e:
        print(e)
        return {'status': False, 'is_private': 0, "id": "00"}


def check_profile(name):
    try:
        users = api.GetUser(screen_name=name)
        return users
    except Exception as e:
        print(e)
        return False
