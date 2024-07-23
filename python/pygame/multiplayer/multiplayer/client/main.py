import configparser
import pygame

from pygame.locals import QUIT, KEYDOWN


def main():
    config = configparser.ConfigParser()
    with open('config.ini', 'r') as configfile:
        config.read_file(configfile)

    H = config.getint('CLIENT', 'H')
    W = config.getint('CLIENT', 'W')
    FPS = config.getint('CLIENT', 'FPS')

    # Init
    pygame.display.init()
    pygame.font.init()
    pygame.key.set_repeat(1, 1)
    pygame.display.set_caption("Multiplayer")
    window = pygame.display.set_mode((W, H))
    clock = pygame.time.Clock()

    while True:
        clock.tick(FPS)

        # Quit event
        events = pygame.event.get()
        if QUIT in [e.type for e in events]:
            break

        # Keyboard events
        if KEYDOWN in [e.type for e in events]:
            keys = pygame.key.get_pressed()
            if keys[pygame.K_ESCAPE]:
                break

            player_dir = [0, 0]

            if keys[pygame.K_w]:
                player_dir[1] -= 1
            if keys[pygame.K_a]:
                player_dir[0] -= 1
            if keys[pygame.K_s]:
                player_dir[1] += 1
            if keys[pygame.K_d]:
                player_dir[0] += 1

            print(player_dir)


if __name__ == "__main__":
    pass
