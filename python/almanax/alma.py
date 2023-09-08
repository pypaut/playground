#!/usr/bin/python3

import datetime
import re
import requests


"""
Almanax quest data retriever for Dofus Touch.
Display items for the next 8 days (including today).
"""

NB_DAYS = 8
URL = "http://www.krosmoz.com/fr/almanax"


def main():
    date = datetime.datetime.today()
    delta = datetime.timedelta(days=1)
    end_date = date + 7 * delta

    while date <= end_date:
        str_date = date.strftime("%Y-%m-%d")
        page = requests.get(f"{URL}/{str_date}").text
        result = re.search(r"Récupérer (.*?) et", page).group(1)
        print(f"{str_date} : {result}")
        date += delta


if __name__ == "__main__":
    main()
