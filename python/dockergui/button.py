import pygame as pg

class Button:
    def __init__(self, text, top, screen_w):
        myfont = pg.font.SysFont("Comic Sans MS", 50)
        self.surface = myfont.render(text, True, (255, 255, 255))
        self.text = text
        w = self.surface.get_width()
        h = self.surface.get_height()
        self.rect = pg.Rect((screen_w - w) / 2, top - h / 2, w, h)

    def draw(self, window):
        window.blit(self.surface, (self.rect[0], self.rect[1]))
