#!/usr/bin/python3

import json
import requests


def main():
    response = requests.get("http://127.0.0.1:8000/foobarqix/3")
    print(response.json())


if __name__ == "__main__":
    main()
