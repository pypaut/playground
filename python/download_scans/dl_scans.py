#!/usr/bin/python3

import fpdf
import os
import requests


def make_dir(name):
    path = os.path.join(name)
    if not os.path.exists(path):
        os.mkdir(path)
    return path


def make_pdf(path, chapter_path, chapter_name):
    pdf_name = f"{chapter_name}.pdf"
    print(f'Build PDF "{pdf_name}"')
    pdf = fpdf.FPDF()
    image_names = sorted(os.listdir(chapter_path))
    for image in image_names:
        image_path = os.path.join(chapter_path, image)
        pdf.add_page()
        pdf.image(image_path, 0, 0, 210, 298)
    pdf_path = os.path.join(path, pdf_name)
    pdf.output(pdf_path, "F")


def download_page(chapter_url, chapter, chapter_path, page):
    page_url = chapter_url + f"{page}.jpg"
    page_url_alt = chapter_url + f"{page}.jpeg"
    print(f"{page_url}... ", end="")
    page_name = str(page).zfill(3)
    filename = f"{chapter}_{page_name}.jpg"
    page_path = os.path.join(chapter_path, filename)
    if not os.path.isfile(page_path):
        r = requests.get(page_url, allow_redirects=True)
        if r.status_code != 200:
            r = requests.get(page_url_alt, allow_redirects=True)
            if r.status_code != 200:
                print("Nothing, end of chapter")
                return False
        print("Download")
        with open(page_path, "wb") as f:
            f.write(r.content)
    else:
        print("Skip")
    return True


def download_chapter(url, chapter, download_path):
    print(f"DOWNLOAD CHAPTER {chapter}")
    chapter_name = str(chapter).zfill(5)
    chapter_path = os.path.join(download_path, chapter_name)
    make_dir(chapter_path)
    chapter_url = url + "1" + str(chapter).zfill(4) + f"000/"
    page = 1

    while True:
        if not download_page(chapter_url, chapter, chapter_path, page):
            break
        page += 1

    make_pdf(download_path, chapter_path, chapter_name)


def main():
    download_folder = make_dir("jujutsu-kaisen")

    url = "https://cdn.kaisenscans.com/file/mangap/2085/"

    chapter = 1
    while True:
        download_chapter(url, chapter, download_folder)
        chapter += 1


if __name__ == "__main__":
    main()
