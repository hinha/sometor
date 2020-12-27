import timeago
import pytz
from datetime import datetime, timedelta

utc = pytz.UTC
nextTime = datetime.now() - datetime.fromtimestamp(1608939754)


# print(aaa)
# input datetime
print(timeago.format(nextTime, datetime.now()))# will print 3 minutes ago


