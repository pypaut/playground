#include <ball.h>

ball *new_ball(float x, float y, float dir_x, float dir_y) {
    ball *b = calloc(1, sizeof(ball));
    b->rect = calloc(1, sizeof(SDL_Rect));

    b->rect->x = x;
    b->rect->y = y;
    b->rect->w = 10;
    b->rect->h = 10;

    b->x = x;
    b->y = y;
    b->dir_x = dir_x;
    b->dir_y = dir_y;
    b->speed = 2.5;

    return b;
}

void destroy_ball(ball *b) {
    free(b->rect);
    free(b);
}

int draw_ball(ball *b, SDL_Renderer *renderer) {
    if (SDL_SetRenderDrawColor(renderer, 200, 100, 200, 255)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        return 1;
    }

    if (SDL_RenderFillRect(renderer, b->rect)) {
        fprintf(stderr, "%s\n", "Error Renderer.FillRect\0");
        return 1;
    }

    return 0;
}

int update_ball(ball *b, int W, int H) {
    // Update position
    float old_x = b->x;
    float old_y = b->y;
    b->x += b->dir_x * b->speed;
    b->y += b->dir_y * b->speed;
    b->rect->x = b->x;
    b->rect->y = b->y;

    // Check wall collision
    if (b->rect->y < 0 || b->rect->y + b->rect->h > H) {
        b->y = old_y;
        b->dir_y = -b->dir_y;
    }

    if (b->rect->x < 0 || b->rect->x + b->rect->w > W) {
        b->x = old_x;
        b->dir_x = -b->dir_x;
    }

    // Check player collision
    // TODO

    // Check other balls collision
    // TODO

    // b->rect->x = b->x;
    // b->rect->y = b->y;

    // Normalize dir vector to speed
    float dir_norm = pow(pow(b->dir_x, 2) + pow(b->dir_y, 2), 0.5);
    b->dir_x = b->dir_x / dir_norm;
    b->dir_y = b->dir_y / dir_norm;

    return 0;
}
