import json
import requests


def get_devices(params={}):
    """
    Return list of devices
    """
    response = requests.get(
        "https://netbox.internal.scaleway.com/api/dcim/devices/?role=network-access-switch",
        params=params,
    )
    devices = response.json()["results"]
    return devices


def display_devices(params={}):
    devices = get_devices(params)
    for device in devices:
        print(f'### {device["display_name"]}')
        print(f'id : {device["id"]}')
        print(f'position : {device["position"]}')
        try:
            print(f'rack : {device["rack"].get("display_name")}')
        except:
            pass
        try:
            print(f'primary ip4 : {device["primary_ip4"]["address"]}')
        except:
            pass
        try:
            print(f'primary ip6 : {device["primary_ip6"]["address"]}')
        except:
            pass
        try:
            print(f'status value : {device["status"]["value"]}')
        except:
            pass
        print("")
