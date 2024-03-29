#include <ball.h>

ball *new_ball(float x, float y, int w, int h) {
    ball *ba = calloc(1, sizeof(ball));
    ba->speed = 5;
    ba->rect = calloc(1, sizeof(SDL_Rect));
    ba->rect->x = x;
    ba->rect->y = y;
    ba->rect->w = w;
    ba->rect->h = h;
    ba->x = x;
    ba->y = y;
    return ba;
}

void destroy_ball(ball *b) {
    free(b->rect);
    free(b);
}

void set_ball_color(ball *ba, int r, int g, int b, int a) {
    ba->r = r;
    ba->g = g;
    ba->b = b;
    ba->a = a;
}

int draw_ball(ball *b, SDL_Renderer *renderer) {
    if (SDL_SetRenderDrawColor(renderer, b->r, b->g, b->b, b->a)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        return 1;
    }

    if (SDL_RenderFillRect(renderer, b->rect)) {
        fprintf(stderr, "%s\n", "Error Renderer.FillRect\0");
        return 1;
    }

    return 0;
}

int update_ball(ball *b, player *p1, player *p2, int H, int W) {
    // Update position
    float old_x = b->x;
    float old_y = b->y;
    b->x += b->dir_x;
    b->y += b->dir_y;
    b->rect->x = b->x;
    b->rect->y = b->y;

    // Handle player collision
    SDL_bool p1_collision = SDL_HasIntersection(b->rect, p1->rect);
    SDL_bool p2_collision = SDL_HasIntersection(b->rect, p2->rect);
    if (p1_collision || p2_collision) {
        b->rect->x = old_x;
        b->rect->y = old_y;
        b->dir_x = -b->dir_x;

        int ball_middle = b->rect->y + b->rect->h/2;
        int player_middle;

        if (p1_collision) {
            player_middle = p1->rect->y + p1->rect->h/2;
        }

        if (p2_collision) {
            player_middle = p2->rect->y + p2->rect->h/2;
        }

        float deviation = 0.1 * (ball_middle - player_middle);
        b->dir_y = deviation;
    }

    // Handle top/bot walls collision
    if (b->rect->y < 0 || b->rect->y + b->rect->h > H) {
        b->rect->y = old_y;
        b->dir_y = -b->dir_y;
    }

    // Handle left/right walls collision
    if (b->rect->x < 0 || b->rect->x + b->rect->w > W) {
        return 1;
    }

    // Normalize dir vector to speed
    float dir_norm = pow(pow(b->dir_x, 2) + pow(b->dir_y, 2), 0.5);
    b->dir_x = b->dir_x * b->speed / dir_norm;
    b->dir_y = b->dir_y * b->speed / dir_norm;

    return 0;
}
