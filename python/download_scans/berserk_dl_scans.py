#!/usr/bin/python

import os
import requests

from bs4 import BeautifulSoup
from PIL import Image

URL = "https://readberserk.com/chapter/berserk-chapter"

# Chapters: 1-374
special_chapters = [
        # "364.5",
        # "099.005",
    "P0",
    "O0",
    "N0",
    "M0",
    "L0",
    "K0",
    "J0",
    "I0",
    "H0",
    "G0",
    "F0",
    "E0",
    "D0",
    "C0",
    "B0",
    "A0",
]


def make_dir(name):
    path = os.path.join(name)
    if not os.path.exists(path):
        os.mkdir(path)
    return path


def get_img_links_from_page(url, chapter_name):
    print("Get chapter")
    response = requests.get(f"{URL}-{chapter_name}")
    soup = BeautifulSoup(response.content, "html.parser")

    print("Get image links")
    img_links = soup.find_all("img", {"class": "pages__img"})
    img_links = [i["src"] for i in img_links]
    img_links = [
        i.replace("\r", "")
        for i in img_links
        # if i.startswith("https://cdn.readberserk.com")  # for certain pages
    ]
    return img_links


def download_images(image_links, chapter_path):
    print("Download images")
    for i, link in enumerate(image_links):
        data = requests.get(link).content
        file_name = str(i).zfill(3) + ".jpg"
        with open(f"{chapter_path}/{file_name}", "wb") as handler:
            handler.write(data)


def build_pdf(chapter_path, pdf_name):
    print("Build PDF")
    image_names = sorted(os.listdir(chapter_path))
    image_paths = [os.path.join(chapter_path, i) for i in image_names]
    images = [Image.open(i).convert("RGB") for i in image_paths]
    images[0].save(
        os.path.join("berserk", pdf_name),
        save_all=True,
        append_images=images[1:],
    )
    return image_paths


def remove_images(image_paths, chapter_path):
    print("Remove images")
    for i in image_paths:
        os.remove(i)
    os.rmdir(chapter_path)


def main():
    path = make_dir("berserk")

    # chapter_names = [str(n).zfill(3) for n in range(192, 375)]
    # chapter_names += special_chapters
    chapter_names = special_chapters

    for chapter_name in chapter_names:
        print(f"Chapter: {chapter_name}")
        chapter_path = make_dir(f"berserk/{chapter_name}")

        image_links = get_img_links_from_page(URL, chapter_name)
        download_images(image_links, chapter_path)

        pdf_name = f"{chapter_name}.pdf"
        image_paths = build_pdf(chapter_path, pdf_name)

        remove_images(image_paths, chapter_path)


if __name__ == "__main__":
    main()
