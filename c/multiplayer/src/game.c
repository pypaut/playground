#include "math.h"

#include <game.h>


game *init_game(int W, int H) {
    SDL_Init(SDL_INIT_VIDEO|SDL_INIT_AUDIO);

    game *g = calloc(1, sizeof(game));

    g->H = H;
    g->W = W;

    g->window = calloc(1, 8);
    g->renderer = calloc(1, 8);

    g->c = new_clock();

    g->is_running = 1;
    g->error = 0;
    g->has_started = 0;

    if (SDL_CreateWindowAndRenderer(W, H, 0, g->window, g->renderer)) {
        fprintf(stderr, "%s\n", "Error on Window/Renderer creation\0");
        return NULL;
    }

    g->p1 = create_player();

    return g;
}

void update_game(game *g) {
    const Uint64 dt = tick(g->c);
    const Uint8 *keys = SDL_GetKeyboardState(NULL);

    float dir_x = 0;
    float dir_y = 0;

    // Go right
    if (keys[SDL_SCANCODE_D] && g->p1->rect->x + g->p1->rect->w < g->W) {
        dir_x++;
    }

    // Go left
    if (keys[SDL_SCANCODE_A] && g->p1->rect->x > 0) {
        dir_x--;
    }

    // Go down
    if (keys[SDL_SCANCODE_S] && g->p1->rect->y + g->p1->rect->h < g->H) {
        dir_y++;
    }

    // Go up
    if (keys[SDL_SCANCODE_W] && g->p1->rect->y > 0) {
        dir_y--;
    }

    // Normalize direction
    float norm = sqrt(pow(dir_x, 2) + pow(dir_y, 2));
    if (norm != 0) {
        dir_x = dir_x / norm;
        dir_y = dir_y / norm;
    }

    // Apply direction
    g->p1->x = g->p1->x + dir_x * (float)dt;
    g->p1->y = g->p1->y + dir_y * (float)dt;

    // Update rect
    update_player_rect(g->p1);
}

void draw_game(game *g) {
    // Background
    if (SDL_SetRenderDrawColor(*g->renderer, 0, 0, 0, 255)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        g->error = 1;
    }

    if (SDL_RenderClear(*g->renderer)) {
        fprintf(stderr, "%s\n", "Error Renderer.RenderClear\0");
        g->error = 1;
    }

    // Player 1
    if (SDL_SetRenderDrawColor(
        *g->renderer,
        g->p1->color->r,
        g->p1->color->g,
        g->p1->color->b,
        g->p1->color->a)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        g->error = 1;
    }

    SDL_RenderFillRect(*g->renderer, g->p1->rect);

    // Final render
    SDL_RenderPresent(*g->renderer);
}

void handle_quit_event(game *g) {
    SDL_Event event;
    while (SDL_PollEvent(&event)) {
        if (event.type == SDL_QUIT) {
            g->is_running = 0;
        }
    }
}

void destroy_game(game *g) {
    free(g->window);
    free(g->renderer);
    destroy_clock(g->c);
    destroy_player(g->p1);
    free(g);

    SDL_Quit();
}
