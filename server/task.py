import logging
import time
import timeago

import pandas as pd
from datetime import datetime, timedelta
from celery import Celery
from pony.orm import Database
from decouple import config
from celery.signals import after_setup_task_logger
from celery.app.log import TaskFormatter

from orakarik.scrape.tweet_scrape import SnTweetScrape
from orakarik.scrape.ig_scrape import SnInstagramScraper
# from . import summary
from summary import Summary

app = Celery(
    'tasks',
    broker=config("URI_REDIS_HOST"),
    backend=config("URI_REDIS_HOST")
)


@after_setup_task_logger.connect
def setup_task_logger(logger, *args, **kwargs):
    for handler in logger.handlers:
        handler.setFormatter(TaskFormatter('%(asctime)s - %(task_id)s - %(task_name)s - %(name)s - %(levelname)s - %('
                                           'message)s'))


db = Database()
db.bind(**{
    'provider': 'mysql',
    "user": config("MYSQL_USERNAME"),
    "passwd": config("MYSQL_PASSWORD"),
    "host": config("MYSQL_HOST"),
    "port": 3306,
    "db": config("MYSQL_DATABASE")
})

app.conf.update(
    CELERY_TASK_SERIALIZER='json',
    CELERY_ACCEPT_CONTENT=['json'],  # Ignore other content
    CELERY_RESULT_SERIALIZER='json',
    CELERY_ENABLE_UTC=True,
    CELERY_TASK_PROTOCOL=1,
    CELERY_TASK_RESULT_EXPIRES=60 * 5,
    MYSQL_DATABASE=db
)

logger = logging.getLogger(__name__)


@app.task
def twitter_scrape_v1(dataSequence):

    task_request = {}
    if isinstance(dataSequence, list):
        if len(dataSequence) > 0 and len(dataSequence) == 1:
            task_request = dataSequence[0]
        else:
            task_request = dataSequence
    else:
        raise Exception("length of dataSequence")

    if task_request["type"] == "account":
        since = datetime.now() - timedelta(days=7)
        until = datetime.now()
    elif task_request["type"] == "hashtag":
        since = datetime.now() - timedelta(days=30)
        until = datetime.now()

    scrape = SnTweetScrape(since.strftime('%Y-%m-%d'), until.strftime('%Y-%m-%d'), 130, proxy=False, proxy_dict={})
    twitter_data = []
    if task_request["type"] == "account":
        twitter_data = scrape.tweetAccount(task_request["keyword"].replace("@", ""), lang="id")
    elif task_request["type"] == "hashtag":
        twitter_data = scrape.tweetHashtag(task_request["keyword"], lang="id")

        df = pd.DataFrame(twitter_data)

        oDB = app.conf.get("MYSQL_DATABASE")
        scumm = Summary(oDB)
        most_egg = scumm.most_engg_user(df)
        if most_egg:
            for x in most_egg:
                result = scumm.find_user_most_engaged(x, task_request["media"], task_request["id"])
                if not result:
                    scumm.insert_most_engaged(x, task_request["media"], task_request["id"])
    else:
        twitter_data = scrape.tweetSearch(task_request, lang="id")

    dataTList = []
    for tw in twitter_data:
        nextTime = datetime.now() - datetime.fromtimestamp(tw['timestamp'])
        tw["str_updated_date"] = timeago.format(nextTime, datetime.now())
        dataTList.append(tw)

    dataTList.sort(key=lambda k: k['timestamp'], reverse=True)
    return {"results": twitter_data, "last_update": dataTList[0]["created_at"]}


@app.task
def instagram_scrape_v1(dataSequence):

    task_request = {}
    if isinstance(dataSequence, list):
        if len(dataSequence) > 0 and len(dataSequence) == 1:
            task_request = dataSequence[0]
        else:
            task_request = dataSequence
    else:
        raise Exception("length of dataSequence")

    if task_request["type"] == "account":
        since = datetime.now() - timedelta(days=7)
        until = datetime.now()
    elif task_request["type"] == "hashtag":
        since = datetime.now() - timedelta(days=30)
        until = datetime.now()

    scrape = SnInstagramScraper(since.strftime('%Y-%m-%d'), until.strftime('%Y-%m-%d'), 130, proxy=False, proxy_dict={})
    ig_data = []
    if task_request["type"] == "account":
        ig_data = scrape.account(task_request["keyword"])
    else:
        ig_data = scrape.hashtag(task_request["keyword"])

    dataTList = []
    for tw in ig_data:
        nextTime = datetime.now() - datetime.fromtimestamp(tw['timestamp'])
        tw["str_updated_date"] = timeago.format(nextTime, datetime.now())
        dataTList.append(tw)

    dataTList.sort(key=lambda k: k['timestamp'], reverse=True)
    return {"results": ig_data, "last_update": dataTList[0]["created_at"]}

@app.task
def add_reflect(a, b):
    return a + b
