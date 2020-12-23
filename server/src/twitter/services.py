import grpc
from google.protobuf.json_format import ParseDict, MessageToDict

from src import twitter_pb2 as pb
from src import twitter_pb2_grpc as pb2_grpc


class Rpc:

    def __init__(self, host, port):
        self.port = port

        if host == "" or host is None:
            self.host = "0.0.0.0"
        else:
            self.host = host

    def Stream(self, payload: dict):
        channel = grpc.insecure_channel(f"{self.host}:{self.port}", options=(('grpc.enable_http_proxy', 0),))
        stub = pb2_grpc.twitterStub(channel)

        request = ParseDict(payload, pb.twitterRequest())
        return MessageToDict(stub.StreamRequest(request))
