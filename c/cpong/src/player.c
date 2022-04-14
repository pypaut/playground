#include <player.h>

player *new_player(int x, int y, int w, int h) {
    player *p = calloc(1, 8);
    p->rect = calloc(1, 8);
    p->rect->x = x;
    p->rect->y = y;
    p->rect->w = w;
    p->rect->h = h;
    return p;
}

void destroy_player(player *p) {
    free(p->rect);
    free(p);
}

void set_player_color(player *p, int r, int g, int b, int a) {
    p->r = r;
    p->g = g;
    p->b = b;
    p->a = a;
}

int draw_player(player *p, SDL_Renderer *renderer) {
    if (SDL_SetRenderDrawColor(renderer, p->r, p->g, p->b, p->a)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        return 1;
    }

    if (SDL_RenderFillRect(renderer, p->rect)) {
        fprintf(stderr, "%s\n", "Error Renderer.FillRect\0");
        return 1;
    }

    return 0;
}
