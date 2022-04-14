#include <ball.h>

ball *new_ball(int x, int y, int w, int h) {
    ball *ba = calloc(1, 8);
    ba->rect = calloc(1, 8);
    ba->rect->x = x;
    ba->rect->y = y;
    ba->rect->w = w;
    ba->rect->h = h;
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
