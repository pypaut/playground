import pygame

from ball import Ball
from game import Game
from player import Player, ControlType


def main():
    game = Game(1000, 800)
    player1 = Player(
        pygame.Rect(100, game.H / 2 - 50, 10, 100),
        pygame.Color(150, 0, 150),
        ControlType.PLAYER1,
    )
    player2 = Player(
        pygame.Rect(game.W - 100, game.H / 2 - 50, 10, 100),
        pygame.Color(150, 0, 150),
        ControlType.PLAYER2,
    )
    ball = Ball(
        pygame.Rect(game.W / 2 - 5, game.H / 2 - 5, 10, 10),
        pygame.Color(255, 255, 255),
        pygame.Rect(0, 0, game.W, game.H),
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

        if not game.is_running and keys[pygame.K_SPACE]:
            ball.dir_x = ball.speed
            game.is_running = True

        # Update
        dt = game.clock.tick(60)
        player1.update(keys, dt, game.W, game.H)
        player2.update(keys, dt, game.W, game.H)
        if not ball.update(player1, player2, dt):
            return

        # Draw
        game.draw(player1, player2, ball)


if __name__ == "__main__":
    main()
