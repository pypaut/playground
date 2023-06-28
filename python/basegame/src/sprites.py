import os
import pygame

from src.constants import (
    ASSETS_PLAYER_PATH,
    PLAYER_HEIGHT,
    PLAYER_SPRITE_HEIGHT_SCALE,
    PLAYER_SPRITE_WIDTH_SCALE,
    PLAYER_WIDTH,
)


def load_sprites_directory(path):
    filenames = sorted(os.listdir(path))
    sprites = [pygame.image.load(os.path.join(path, f)) for f in filenames]
    sprites = scaled_sprites(sprites)
    return sprites


def flipped_sprites(sprites):
    return [pygame.transform.flip(s, True, False) for s in sprites]


def scaled_sprites(sprites):
    x_scale = PLAYER_WIDTH * PLAYER_SPRITE_WIDTH_SCALE
    y_scale = PLAYER_HEIGHT * PLAYER_SPRITE_HEIGHT_SCALE
    return [pygame.transform.scale(s, [x_scale, y_scale]) for s in sprites]


def load_player_sprites(animations):
    sprites = {
        "right": {
            a: load_sprites_directory(os.path.join(ASSETS_PLAYER_PATH, a))
            for a in animations
        },
    }
    sprites["left"] = {
        a: flipped_sprites(s) for (a, s) in sprites["right"].items()
    }
    return sprites
