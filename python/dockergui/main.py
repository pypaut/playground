#!/usr/bin/python3

import docker
import pygame as pg

from button import Button
from pygame.locals import QUIT, KEYDOWN, MOUSEBUTTONUP

W = 1000
H = 800
FPS = 60.0


def draw_buttons(buttons, window):
    window.fill((0, 0, 0))
    for b in buttons:
        b.draw(window)
    pg.display.flip()


def get_details_buttons(name):
    client = docker.from_env()
    container = client.containers.get(name)
    return [
        Button(f"Status: {container.status}", 150, W),
        Button(f"ID: {container.id}", 200, W),
        Button(f"Image: {container.image}", 250, W),
    ]


def container_details_menu(name, window, clock):
    button_back = Button("BACK", H - 50, W)
    button_title = Button(name, 50, W)

    details_buttons = get_details_buttons(name)

    while True:
        # Events
        events = pg.event.get()
        events = [e.type for e in events]
        if QUIT in events:
            return exit(0)

        # Click event
        if MOUSEBUTTONUP in events:
            mouse_x, mouse_y = pg.mouse.get_pos()
            if button_back.rect.collidepoint(mouse_x, mouse_y):
                return

        # Update
        dt = clock.tick(FPS)

        # Draw
        draw_buttons([button_title, button_back] + details_buttons, window)


def get_containers_buttons():
    client = docker.from_env()
    containers = client.containers.list()
    containers_names = list(map(lambda x: x.name, containers))
    buttons = [Button(containers_names[i], 50 * (i + 1), W) for i in range(len(containers))]
    return buttons


def containers_menu(window, clock):
    button_back = Button("BACK", H - 50, W)
    buttons_containers = get_containers_buttons()
    while True:
        # Events
        events = pg.event.get()
        events = [e.type for e in events]
        if QUIT in events:
            return exit(0)

        # Click event
        if MOUSEBUTTONUP in events:
            mouse_x, mouse_y = pg.mouse.get_pos()
            if button_back.rect.collidepoint(mouse_x, mouse_y):
                return
            for b in buttons_containers:
                if b.rect.collidepoint(mouse_x, mouse_y):
                    container_details_menu(b.text, window, clock)

        # Update
        dt = clock.tick(FPS)

        # Draw
        draw_buttons(buttons_containers + [button_back], window)


def main():
    pg.display.init()
    pg.font.init()
    pg.display.set_caption("Docker GUI")
    window = pg.display.set_mode((W, H))
    clock = pg.time.Clock()

    button_containers = Button("CONTAINERS", H / 2, W)
    button_quit = Button("QUIT", H - 50, W)

    while True:
        # Events
        events = pg.event.get()
        events = [e.type for e in events]
        if QUIT in events:
            return exit(0)

        # Click event
        if MOUSEBUTTONUP in events:
            mouse_x, mouse_y = pg.mouse.get_pos()
            if button_containers.rect.collidepoint(mouse_x, mouse_y):
                containers_menu(window, clock)
            if button_quit.rect.collidepoint(mouse_x, mouse_y):
                return exit(0)

        # Update
        dt = clock.tick(FPS)

        # Draw
        draw_buttons([button_containers, button_quit], window)


if __name__ == "__main__":
    main()
