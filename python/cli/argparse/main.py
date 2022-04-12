#!/usr/bin/python3

import netbox
import parser


def main():
    params = parser.parse()

    # devices = netbox.get_devices(params)
    devices = netbox.display_devices(params)


if __name__ == "__main__":
    main()
