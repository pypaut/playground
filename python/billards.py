#!/usr/bin/python3

import pygame
import pymunk

from pygame.locals import QUIT
from pymunk import Vec2d
from numpy.linalg import norm
from random import randint

# Physics collision types
COLLTYPE_DEFAULT = 0
COLLTYPE_MOUSE = 1
COLLTYPE_BALL = 2

# Constants
W = 1000
H = 800
FPS = 60.0
MAX_STRENGTH = 500
MULT_FACTOR = 10


def flipy(y):
    """
    Small hack to convert chipmunk physics to pygame coordinates
    """
    return -y + H


def create_ball(x, y):
    """
    Create body and shape for a new ball at position (x, y)
    (PyGame coordinates)
    """
    body = pymunk.Body(1, 1)
    body.position = (x, flipy(y))
    shape = pymunk.Circle(body, 10, (0, 0))
    shape.friction = 0.9
    shape.elasticity = 0.8
    shape.collision_type = COLLTYPE_BALL
    return body, shape


def create_wall(x1, y1, x2, y2, space):
    """
    Create segment from (x1, y1) to (x2, y2), shape only
    """
    shape = pymunk.Segment(
        space.static_body, (x1, flipy(y1)), (x2, flipy(y2)), 0.0
    )
    shape.friction = 0.99
    shape.elasticity = 0.8
    return shape


def draw_ball(window, shape, color):
    """
    Draw ball shape to screen
    """
    r = shape.radius
    v = shape.body.position
    rot = shape.body.rotation_vector
    p = int(v.x), int(flipy(v.y))
    p2 = p + Vec2d(rot.x, -rot.y) * r * 0.9
    p2 = int(p2.x), int(p2.y)
    pygame.draw.circle(window, color, p, int(r), 2)
    pygame.draw.line(window, pygame.Color("red"), p, p2)


def draw_strength(window, shape, charging_direction):
    """
    Draw direction line to measure strike strength
    """
    v = shape.body.position
    p1 = int(v.x), int(flipy(v.y))
    p2 = p1 + Vec2d(charging_direction[0], -charging_direction[1]) // 4
    pygame.draw.line(window, pygame.Color("white"), p1, p2)


def main():
    # Init
    pygame.display.init()
    pygame.display.set_caption("PyMunk Experiment")
    window = pygame.display.set_mode((W, H))

    # PyMunk physics
    space = pymunk.Space()
    space.gravity = 0.0, 0.0

    # Ball init
    body, shape = create_ball(W // 2, H // 2)
    space.add(body, shape)
    colors = [pygame.Color("white")]

    # Other balls
    balls = [shape]
    for i in range(30):
        new_x = randint(10, W - 10)
        new_y = randint(10, H - 10)
        new_body, new_shape = create_ball(new_x, new_y)
        balls.append(new_shape)
        colors.append(pygame.Color("blue"))
        space.add(new_body, new_shape)

    # Walls
    walls = [
        create_wall(0, flipy(H), W, flipy(H), space),
        create_wall(0, flipy(0), W, flipy(0), space),
        create_wall(0, flipy(H), 0, flipy(0), space),
        create_wall(W, flipy(H), W, flipy(1), space),
    ]
    for w in walls:
        space.add(w)

    # State variables
    charging_strike = False
    charging_direction = (0, 0)
    click_position = (0, 0)

    while True:
        # Events
        events = pygame.event.get()
        if QUIT in [e.type for e in events]:
            break

        mouse_pressed = pygame.mouse.get_pressed(num_buttons=3)[0]
        if mouse_pressed:
            mouse_pos = pygame.mouse.get_pos()
            if charging_strike:
                x_charging = -(mouse_pos[0] - click_position[0]) * MULT_FACTOR
                y_charging = (mouse_pos[1] - click_position[1]) * MULT_FACTOR
                charging_direction = Vec2d(x_charging, y_charging)
                n = norm(charging_direction, 2)
                if n > MAX_STRENGTH:
                    charging_direction *= MAX_STRENGTH / n
            else:
                charging_strike = True
                click_position = mouse_pos
                charging_direction = (0, 0)
        else:
            if charging_strike:
                charging_strike = False
                print(f"Applied {charging_direction}")
                body.apply_impulse_at_world_point(
                    charging_direction, point=shape.body.position
                )
                charging_direction = (0, 0)

        # Update
        space.step(1.0 / FPS)

        # Draw
        window.fill((0, 0, 0))
        for i in range(len(balls)):
            draw_ball(window, balls[i], colors[i])
        if charging_strike:
            draw_strength(window, shape, charging_direction)
            pygame.draw.circle(
                window, pygame.Color("red"), click_position, radius=5
            )
        pygame.display.flip()

    pygame.quit()


if __name__ == "__main__":
    main()
