import argparse
from src import server_twitter, server_instagram
from src.api import rest_api


def parse_args():
    parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    parser.add_argument('-i', "--init", dest="init")
    parser.add_argument("-w", "--worker", dest="worker")
    parser.add_argument('-n', '--service-name', dest='service_name',
                        type=lambda x: str(x) if str(x) != '' else parser.error('--service-name N must be string'),
                        metavar='N',
                        help='Only return the first N results')
    group = parser.add_mutually_exclusive_group(required=False)
    group.add_argument('-p', '--port', dest='port',
                       type=lambda x: int(x) if int(x) > 1 else parser.error('--port N must required'),
                       metavar='N', help='Port')

    args = parser.parse_args()

    if not args.port:
        raise RuntimeError('Error: no port specified')

    return args


def main():
    args = parse_args()
    if args.service_name == 'twitter':
        server_twitter.serve(args.port)
    if args.service_name == "instagram":
        server_instagram.serve(args.port)
    if args.init == "twitter":
        if not args.worker:
            args.worker = 1
        rest_api(args.port, int(args.worker))


# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    main()
