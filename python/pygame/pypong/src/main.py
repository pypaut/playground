import pygame

from game import Game
from player import Player, ControlType


def main():
    game = Game(1000, 800)

    player1 = Player(
        pygame.Rect(100, game.H / 2 - 50, 10, 100),
        ControlType.PLAYER1,
    )
    player2 = Player(
        pygame.Rect(game.W - 100, game.H / 2 - 50, 10, 100),
        ControlType.PLAYER2,
    )

    while True:
        # Events
        events = pygame.event.get()
        events_types = [e.type for e in events]
        if pygame.QUIT in events_types:
            return

        keys = pygame.key.get_pressed()
        if keys[pygame.K_ESCAPE]:
            return

        # Update
        dt = game.clock.tick(60)
        player1.update(keys, dt, game.W, game.H)
        player2.update(keys, dt, game.W, game.H)

        # Draw
        game.window.fill((0, 0, 0))
        pygame.draw.rect(game.window, (150, 0, 150), player1.rect)
        pygame.draw.rect(game.window, (150, 0, 150), player2.rect)
        pygame.display.flip()


if __name__ == "__main__":
    main()
