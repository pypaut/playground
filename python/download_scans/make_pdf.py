import os

from PIL import Image

chapter_name = "001"
chapter_path = "berserk/001"

pdf_name = "001.pdf"

image_names = sorted(os.listdir(chapter_path))
images = [Image.open(os.path.join(chapter_path, i)) for i in image_names]
images = [i.convert('RGB') for i in images]

images[0].save(os.path.join(chapter_path, pdf_name), save_all=True, append_images=images[1:])

# for image_name in image_names:
#     image_path = os.path.join(chapter_path, image_name)
#     pdf.add_page()
#     pdf.image(image_path, 0, 0, 210, 298)
# pdf_path = os.path.join(chapter_path, pdf_name)
# pdf.output(pdf_path, "F")
