from PIL import Image

a = np.zeros((500, 500, 3))
im = Image.fromarray(a)
im.save("pic.png")
