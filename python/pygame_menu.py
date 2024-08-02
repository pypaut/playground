#!/usr/bin/python3

import pygame

from pygame.locals import QUIT, KEYDOWN, MOUSEBUTTONDOWN


W = 1000
H = 800
FPS = 60.0

TRANSIT_SURF = pygame.Surface((W, H))
TRANSIT_SURF.fill((0, 0, 0))
TRANSIT_SPEED = 10


def screen_transition(window, clock):
    """ """
    for i in range(0, 100, TRANSIT_SPEED):
        clock.tick(FPS)
        TRANSIT_SURF.set_alpha(i)
        window.blit(TRANSIT_SURF, (0, 0))
        pygame.display.flip()


def menu_title_surface(title):
    """
    Instanciate surface for title
    """
    title_font = pygame.font.SysFont("Comic Sans MS", 80)
    title_surface = title_font.render(title, True, (255, 255, 255))
    title_w = title_surface.get_width()
    title_h = title_surface.get_height()
    title_pos = (W / 2 - title_w / 2, H / 6 - title_h)
    return title_surface, title_pos


def generate_buttons_rects(buttons):
    """
    Compute list of rects for each button
    """
    buttons_left = W / 4
    buttons_w = W / 2
    buttons_h = H / 10
    gap = buttons_h / 3
    buttons_h -= gap
    buttons_tops = [H / 3 + i * (buttons_h + gap) for i in range(len(buttons))]
    rects = [
        pygame.Rect(buttons_left, top, buttons_w, buttons_h)
        for top in buttons_tops
    ]
    return rects


def generate_buttons_surfaces(buttons, buttons_rects):
    """
    Instantiate text surface for each button
    """
    button_font = pygame.font.SysFont("Comic Sans MS", 40)
    buttons_surfaces = [
        button_font.render(b[0], True, (0, 0, 0)) for b in buttons
    ]

    buttons_pos = []
    for i in range(len(buttons)):
        rect = buttons_rects[i]
        surf = buttons_surfaces[i]
        button_w = W / 2 - buttons_surfaces[i].get_width() / 2
        button_h = rect[1] + surf.get_height() / 2
        buttons_pos.append((button_w, button_h))

    return buttons_surfaces, buttons_pos


def check_selected(mouse_pos, buttons_rects):
    """
    Check if mouse_pos collides with any button
    """
    is_selected = False
    count = 0
    for b in buttons_rects:
        if b.collidepoint(mouse_pos):
            is_selected = True
            break
        count += 1
    if is_selected:
        return count
    return None


def menu(
    window, clock, title="MENU", buttons=[], back_with_esc=True, anim=True
):
    """
    Display new menu on @window, with title @title and displayed @buttons

    buttons : [(str, str)] list of form [(title, signal)]
    """
    # Transition animation
    if anim:
        screen_transition(window, clock)

    # Title text
    title_surface, title_pos = menu_title_surface(title)

    # Buttons rects
    buttons_rects = generate_buttons_rects(buttons)

    # Buttons text surfaces
    buttons_surfaces, buttons_pos = generate_buttons_surfaces(
        buttons, buttons_rects
    )

    # Main loop variables
    index_selected = None
    alpha = 0

    # Main loop
    while True:
        # Check mouse pos
        mouse_pos = pygame.mouse.get_pos()
        index_selected = check_selected(mouse_pos, buttons_rects)

        # Events
        for e in pygame.event.get():
            if index_selected != None and e.type == MOUSEBUTTONDOWN:
                return buttons[index_selected][1]
            if back_with_esc and e.type == KEYDOWN:
                keys = pygame.key.get_pressed()
                if keys[pygame.K_ESCAPE]:
                    return "back"

        # Update
        dt = clock.tick(FPS)

        # Display
        window.fill((0, 0, 0))
        window.blit(title_surface, title_pos)
        for i in range(len(buttons)):
            color = (200, 200, 200)
            if i == index_selected:
                color = (255, 255, 255)
            pygame.draw.rect(window, color, buttons_rects[i])
            window.blit(buttons_surfaces[i], buttons_pos[i])
        if anim and alpha < 100:
            TRANSIT_SURF.set_alpha(100 - alpha)
            window.blit(TRANSIT_SURF, (0, 0))
            alpha += TRANSIT_SPEED
        pygame.display.flip()


def game(window, clock):
    """
    Game loop, can pause
    """
    myfont = pygame.font.SysFont("Comic Sans MS", 60)
    text_surface = myfont.render("GAME", True, (255, 255, 255))
    text_w = text_surface.get_width()
    text_pos = (W / 2 - text_w / 2, H / 2)

    while True:
        # Events
        events = pygame.event.get()
        if QUIT in [e.type for e in events]:
            return exit(0)

        # Keyboard events
        if KEYDOWN in [e.type for e in events]:
            keys = pygame.key.get_pressed()
            if keys[pygame.K_ESCAPE]:
                # command = pause(window, clock)
                command = menu(
                    window,
                    clock,
                    "PAUSE",
                    [
                        ("Continue", "continue"),
                        ("Retry", "retry"),
                        ("Quit", "quit"),
                    ],
                    back_with_esc=True,
                    anim=False,
                )
                if command == "quit":
                    return
                if command == "restart":
                    return  # FIXME

        # Update
        dt = clock.tick(FPS)

        # Display
        window.fill((0, 0, 0))
        window.blit(text_surface, text_pos)
        pygame.display.flip()


def main():
    # Init
    pygame.display.init()
    pygame.font.init()
    pygame.display.set_caption("PyGame basic template")
    window = pygame.display.set_mode((W, H))
    clock = pygame.time.Clock()

    # Main loop
    while True:
        # command = main_menu(window, clock)
        command = menu(
            window,
            clock,
            "MAIN MENU",
            [("Start", "start"), ("Options", "options"), ("Quit", "back")],
        )

        if command == "back":
            break  # TODO : Alert menu "Are you sure?"

        elif command == "start":
            game(window, clock)

        else:  # command == "options"
            menu(
                window,
                clock,
                "OPTIONS",
                [
                    ("Option 1", ""),
                    ("Option 2", ""),
                    ("Option 3", ""),
                    ("Back", "back"),
                ],
            )


if __name__ == "__main__":
    main()
