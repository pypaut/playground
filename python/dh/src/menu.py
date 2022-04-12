import pygame

from pygame.locals import KEYDOWN
from src.constants import W, H, FPS


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
    Instantiate surface for title
    """
    title_font = pygame.font.SysFont("Comic Sans MS", 80)
    title_surface = title_font.render(title, True, (255, 255, 255))
    title_w = title_surface.get_width()
    title_h = title_surface.get_height()
    title_pos = (W / 2 - title_w / 2, H / 3 - title_h)
    return title_surface, title_pos


def generate_buttons_rects(buttons):
    """
    Compute list of rects for each button
    """
    buttons_w = W / 5
    buttons_left = (W - buttons_w) / 2
    buttons_h = (H / 5) / len(buttons)
    gap = buttons_h / 3
    buttons_h -= gap
    buttons_tops = [
        H * 2 / 3 + i * (buttons_h + gap) for i in range(len(buttons))
    ]
    rects = [
        pygame.Rect(buttons_left, top, buttons_w, buttons_h)
        for top in buttons_tops
    ]
    return rects


def generate_buttons_surfaces(buttons):
    """
    Instantiate text surface for each button
    """
    buttons_rects = generate_buttons_rects(buttons)
    button_font = pygame.font.SysFont("Comic Sans MS", 30)
    buttons_surfaces = [
        button_font.render(b[0], True, (255, 255, 255)) for b in buttons
    ]
    buttons_pos = [
        (
            W / 2 - buttons_surfaces[i].get_width() / 2,
            buttons_rects[i][1] + buttons_surfaces[i].get_height(),
        )
        for i in range(len(buttons))
    ]
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


def menu(window, clock, title="MENU", buttons=None, anim=True):
    """
    Display new menu on @window, with title @title and displayed @buttons

    buttons : [(str, str)] list of form [(title, signal)]
    """
    # Transition animation
    if buttons is None:
        buttons = []
    if anim:
        screen_transition(window, clock)

    # Title text
    title_surface, title_pos = menu_title_surface(title)

    # Buttons text surfaces
    buttons_surfaces, buttons_pos = generate_buttons_surfaces(buttons)

    # Cursor surface
    font = pygame.font.SysFont("Comic Sans MS", 30)
    cursor_surf = font.render("> ", True, (255, 255, 255))

    # Main loop variables
    index_selected = 0
    alpha = 0

    # Main loop
    while True:
        # Events
        for e in pygame.event.get():
            if e.type == KEYDOWN:
                keys = pygame.key.get_pressed()
                if keys[pygame.K_DOWN]:
                    index_selected += 1
                elif keys[pygame.K_UP]:
                    index_selected -= 1
                elif keys[pygame.K_RETURN]:
                    return buttons[index_selected][1]
                else:
                    pass
                index_selected %= len(buttons)

        # Update
        clock.tick(FPS)

        # Display
        window.fill((0, 0, 0))
        window.blit(title_surface, title_pos)
        for i in range(len(buttons)):
            window.blit(buttons_surfaces[i], buttons_pos[i])
        x, y = buttons_pos[index_selected]
        window.blit(cursor_surf, (x - 30, y))
        if anim and alpha < 100:
            TRANSIT_SURF.set_alpha(100 - alpha)
            window.blit(TRANSIT_SURF, (0, 0))
            alpha += TRANSIT_SPEED
        pygame.display.flip()


def pause_menu(window, clock):
    """
    Pause menu
    """
    return menu(
        window,
        clock,
        "PAUSE",
        [
            ("Continue", "continue"),
            ("Quit", "quit"),
        ],
        anim=False,
    )
