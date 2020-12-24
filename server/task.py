import logging
import time
from celery import Celery
from pony.orm import Database
from decouple import config
from celery.signals import after_setup_task_logger
from celery.app.log import TaskFormatter

from orakarik.scrape.tweet_scrape import SnTweetScrape

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
from datetime import datetime, timedelta


@app.task
def twitter_scrape_v1(dataSequence):
    # oDB = app.conf.get("MYSQL_DATABASE")
    # keyword = []
    # with db_session:
    #     for data in oDB.execute("select * from stream_sequence_account"):
    #         keyword.append(data[1])

    task_request = {}
    if isinstance(dataSequence, list):
        if len(dataSequence) > 0 and len(dataSequence) == 1:
            task_request = dataSequence[0]
        else:
            task_request = dataSequence
    else:
        raise Exception("length of dataSequence")

    since = datetime.now() - timedelta(2)
    until = datetime.now()

    scrape = SnTweetScrape(since.strftime('%Y-%m-%d'), until.strftime('%Y-%m-%d'), 100, proxy=False, proxy_dict={})
    twitter_data = []
    if task_request["type"] == "account":
        twitter_data = scrape.tweetAccount(task_request["keyword"].replace("@", ""), lang="id")
    elif task_request["type"] == "hashtag":
        twitter_data = scrape.tweetHashtag(task_request, lang="id")
    else:
        twitter_data = scrape.tweetSearch(task_request, lang="id")

    twitter_data = sorted(twitter_data, key=lambda k: k['timestamp'])
    return {"results": twitter_data, "last_update": twitter_data[0]["created_at"]}


@app.task
def add_reflect(a, b):
    return a + b
