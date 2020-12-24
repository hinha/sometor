import json
# import re
# import ast
import pandas as pd
from pony.orm import db_session


class Summary:

    def __init__(self, db):
        self.db = db

    # most active username
    def most_act_user(self, posts):
        df = posts[['user_name', 'screen_name']]
        df1 = df.groupby(['user_name', 'screen_name']).user_name.agg('count').to_frame('count')
        df1 = df1.sort_values(by=['count'], ascending=False).head(10).reset_index()
        df2 = df.drop_duplicates(['user_name'])
        df3 = pd.merge(df1, df2, on=['user_name', 'screen_name'], how='left')

        # return to json
        return json.dumps(df3.to_dict('records'))

    def top_ment_user(self, posts):
        top_mentioned = posts[['mentions']]
        top_mentioned = top_mentioned.dropna()

        username = []
        for x in top_mentioned['mentions']:
            for text in x:
                username.append(text['text'].replace("@", ""))

        top_mentioned = pd.DataFrame({'username': username})
        top_mentioned = top_mentioned['username'].value_counts()
        top_mentioned = pd.DataFrame({'username': list(top_mentioned.index), 'count': list(top_mentioned)})
        top_mentioned = top_mentioned.sort_values(by=['count'], ascending=False).reset_index(drop=True).head(10)

        return json.dumps(top_mentioned.to_dict('records'))

    def most_engg_user(self, posts):
        most_engaged_user = posts[['user_name', 'reply_count', 'retweet_count', 'like_count', 'quote_count']]

        most_engaged_user = most_engaged_user.assign(
            total_engagement=lambda most_engaged:
            most_engaged['reply_count'] + most_engaged['retweet_count'] + most_engaged['like_count'] + most_engaged[
                'quote_count']
        )

        most_engaged_user = most_engaged_user.groupby('user_name', as_index=False) \
            .agg({'total_engagement': "sum"}) \
            .sort_values(by=['total_engagement'], ascending=False) \
            .head(20).reset_index(drop=True)

        df = posts[['user_name']]
        df2 = df.drop_duplicates(['user_name'])
        df3 = pd.merge(most_engaged_user, df2, on=['user_name'], how='left')
        # return df3
        return df3.to_dict('records')

    def insert_most_engaged(self, raw: dict, media: str, accountID: str):
        with db_session:
            sql = "insert into most_engaged_user (user_name,total_engagement,media,created_at,stream_sequence_account_id) " + \
                  "VALUES ('{}',{},'{}',now(),'{}');".format(raw["user_name"], raw["total_engagement"], media, accountID)
            print(sql)
            self.db.execute(sql)
            print("inserted")

    def find_user_most_engaged(self, raw: dict, media: str, accountID: str):
        with db_session:
            sql = f"select * from most_engaged_user where user_name = '{raw['user_name']}' and media = '{media}' and stream_sequence_account_id = '{accountID}'"
            result = self.db.execute(sql)
            data = result.fetchone()
            if data:
                return data
            else:
                return None
# if __name__ == '__main__':
#     obj = Summary()
#     data = pd.read_csv("jokowi.csv")
#     print(obj.most_act_user(data))
