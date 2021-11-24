import argparse
import json
import sys


def parse():
    parser = argparse.ArgumentParser(description="Netbox API for Scaleway")
    parser.add_argument(
        "-p",
        "--params",
        metavar="params",
        type=str,
        help="params for the request, formatted as JSON",
    )
    args = parser.parse_args()

    if args.params == None:
        return None

    try:
        params = json.loads(args.params)
    except json.decoder.JSONDecodeError:
        sys.exit('ERROR : params should look like : \'{"id":"178546"}\'')

    return params
