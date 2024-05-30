#include <player.h>

player *create_player() {
    player *p = calloc(1, sizeof(player));

    p->x = 100;
    p->y = 100;

    p->w = 50;
    p->h = 50;

    p->rect= calloc(1, sizeof(SDL_Rect));
    p->rect->x = (int)p->x;
    p->rect->y = (int)p->y;
    p->rect->w = (int)p->w;
    p->rect->h = (int)p->h;

    p->color = calloc(1, sizeof(SDL_Color));
    p->color->r = 200;
    p->color->g = 200;
    p->color->b = 200;
    p->color->a = 255;

    return p;
}

void destroy_player(player *p) {
    free(p->rect);
    free(p->color);
    free(p);
}

void update_player_rect(player *p) {
    p->rect->x = (int)p->x;
    p->rect->y = (int)p->y;
    p->rect->w = (int)p->w;
    p->rect->h = (int)p->h;
}
