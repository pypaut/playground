#include <player.h>


player *new_player(int H, int W) {
    player *p = calloc(1, sizeof(player));
    p->rect = calloc(1, sizeof(SDL_Rect));

    p->rect->x = W / 2 - 10;
    p->rect->y = H / 2 - 10;
    p->rect->w = 20;
    p->rect->h = 20;

    return p;
}

void destroy_player(player *p) {
    free(p->rect);
    free(p);
}

int draw_player(player *p, SDL_Renderer *renderer) {
    if (SDL_SetRenderDrawColor(renderer, 255, 255, 255, 255)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        return 1;
    }

    if (SDL_RenderFillRect(renderer, p->rect)) {
        fprintf(stderr, "%s\n", "Error Renderer.FillRect\0");
        return 1;
    }

    return 0;
}

void update_player(player *p, int mouse_x, int mouse_y) {
    p->rect->x = mouse_x - p->rect->w / 2;
    p->rect->y = mouse_y - p->rect->h / 2;
}
