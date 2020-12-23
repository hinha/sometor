import grpc
from google.protobuf.json_format import ParseDict

from src import twitter_pb2 as pb
from src import twitter_pb2_grpc as pb2_grpc


def run_twitter():
    channel = grpc.insecure_channel("0.0.0.0:50055")
    stub = pb2_grpc.twitterStub(channel)

    payload = {
        "keyword": "@detikcom",
        "search_type": "account",
        "since": "2020-10-26",
        "until": "2020-10-27"
    }

    request = ParseDict(payload, pb.twitterRequest())
    stub.StreamRequest(request)


def nikParse():
    pass



if __name__ == '__main__':
    run_twitter()
