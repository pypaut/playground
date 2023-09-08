#!/usr/bin/python3

import cv2 as cv
import numpy as np
import sys


global_vars = {"nb": 0}  # Little trick for global vars


def mouse_callback(event, x, y, flags, params):
    """
    Click on the contour you want to save as image
    """
    contours = params[0]
    spritesheet = params[1]
    global_vars = params[2]
    if event == cv.EVENT_LBUTTONDOWN:
        for c in contours:
            r = cv.pointPolygonTest(c, (x, y), False)
            if r > 0:
                b = cv.boundingRect(c)
                extracted = spritesheet[b[1] : b[1] + b[3], b[0] : b[0] + b[2]]
                nb = global_vars["nb"]
                name = f"{nb}.png"
                cv.imwrite(name, extracted)
                print(f"Saved {name}")
                global_vars["nb"] += 1
                continue


def main():
    # Load spritesheet
    try:
        sprite_file = sys.argv[1]
        spritesheet = cv.imread(sprite_file)
    except:
        print(
            "Cannot find your spritesheet. Usage : ./sprite_ripper.py <spritesheet_location>"
        )
        return

    # Detect contours
    gray_spritesheet = cv.cvtColor(spritesheet, cv.COLOR_BGR2GRAY)
    mask = cv.inRange(gray_spritesheet, 0, 250)
    # kernel = cv.getStructuringElement(cv.MORPH_ELLIPSE, (1, 1))
    # mask = cv.morphologyEx(mask, cv.MORPH_CLOSE, kernel)
    contours, hierarchy = cv.findContours(
        mask, cv.RETR_TREE, cv.CHAIN_APPROX_NONE
    )

    # Display contours for selection
    display_spritesheet = spritesheet.copy()
    cv.drawContours(
        display_spritesheet, contours, -1, (0, 0, 255)
    )  # -1 for all contours
    cv.imshow("Contours", display_spritesheet)
    cv.setMouseCallback(
        "Contours", mouse_callback, [contours, spritesheet.copy(), global_vars]
    )
    cv.waitKey(0)


if __name__ == "__main__":
    main()
