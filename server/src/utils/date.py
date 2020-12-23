from datetime import datetime, timedelta


class DateSettings(object):

    def __init__(self, since, until):
        self.since = since
        self.until = until

    def get_date(self):

        mapObj = {}
        if self.__is_week():
            dateStr, limits, lengthDate = self.weeks_between(2)
            if self.__date_overlap(dateStr):
                mapObj['since'] = (datetime.now() - timedelta(days=lengthDate)).strftime("%Y-%m-%d")
                mapObj['until'] = datetime.now().strftime("%Y-%m-%d")
                mapObj['count'] = limits
                mapObj['step'] = datetime.now().strftime("%Y-%m-%d")
                mapObj['length'] = lengthDate
            else:
                mapObj['since'] = (datetime.strptime(dateStr, "%Y-%m-%d") - timedelta(days=lengthDate)).strftime(
                    "%Y-%m-%d")
                mapObj['until'] = datetime.now().strftime("%Y-%m-%d")
                mapObj['count'] = limits
                mapObj['step'] = (datetime.strptime(dateStr, "%Y-%m-%d") - timedelta(days=lengthDate + 1)).strftime(
                    "%Y-%m-%d")
                mapObj['length'] = lengthDate
        else:
            mapObj['since'] = self.since
            mapObj['until'] = self.until
            mapObj['count'] = 200
            mapObj['step'] = self.until
            mapObj['length'] = 1

        return mapObj

    """
        strdate: str    YYYY-mm-dd
        return: bool
        description: hitung rentang waktu dengan tanggal sekarang
    """

    def __date_overlap(self, strdate):
        ts = datetime.now().timestamp()
        overdate = self.__str2date(strdate)

        if ts >= overdate:
            return False
        else:
            return True

    """
        self: mengecek jika rentang waktu lebih 1 minggu
        return: bool

        description: jika kurang dari 1 minggu meta data keyword di set-default
        status akan dibuat menjadi done
    """

    def __is_week(self):
        since = datetime.strptime(self.since, "%Y-%m-%d")
        until = datetime.strptime(self.until, "%Y-%m-%d")
        delta = until - since
        if delta.days > 0 and delta.days >= 7:
            return True
        else:
            return False

    def __str2date(self, strdate):
        date = datetime.strptime(strdate, "%Y-%m-%d")
        timestamp = datetime.timestamp(date)
        return timestamp

    def weeks_between(self, length):
        weeksList = []
        cursor = datetime.strptime(self.since, "%Y-%m-%d")
        cursor_end = datetime.strptime(self.until, "%Y-%m-%d")
        limits = 100

        end = ""
        while cursor <= cursor_end:
            if cursor.weekday() not in weeksList:
                weeksList.append(cursor.day)
                end = cursor
            limits += 50
            cursor += timedelta(weeks=1)

        leasts = cursor - timedelta(weeks=2)

        if len(weeksList) > length:
            return leasts.strftime("%Y-%m-%d"), limits, len(weeksList)
        return end.strftime("%Y-%m-%d"), limits, len(weeksList)