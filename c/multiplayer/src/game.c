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

    // if (SDL_SetWindowFullscreen(*g->window, SDL_WINDOW_FULLSCREEN_DESKTOP)) {
    //     fprintf(stderr, "%s\n", "error on SDL_SetWindowFullscreen()\0");
    //     return NULL;
    // }

    g->p = create_player();

    return g;
}

void update_game(game *g) {
    update_game_p(g);
}

void update_game_p(game *g) {
    const Uint64 dt = tick(g->c);
    const Uint8 *keys = SDL_GetKeyboardState(NULL);

    float dir_x = 0;
    float dir_y = 0;

    // Go right
    if (keys[SDL_SCANCODE_D] && g->p->rect->x + g->p->rect->w < g->W) {
        dir_x++;
    }

    // Go left
    if (keys[SDL_SCANCODE_A] && g->p->rect->x > 0) {
        dir_x--;
    }

    // Go down
    if (keys[SDL_SCANCODE_S] && g->p->rect->y + g->p->rect->h < g->H) {
        dir_y++;
    }

    // Go up
    if (keys[SDL_SCANCODE_W] && g->p->rect->y > 0) {
        dir_y--;
    }

    // Normalize direction
    float norm = sqrt(pow(dir_x, 2) + pow(dir_y, 2));
    if (norm != 0) {
        dir_x = dir_x / norm;
        dir_y = dir_y / norm;
    }

    // Apply direction
    g->p->x = g->p->x + dir_x * (float)dt;
    g->p->y = g->p->y + dir_y * (float)dt;

    // Update rect
    update_player_rect(g->p);
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

    // Player
    if (SDL_SetRenderDrawColor(
        *g->renderer,
        g->p->color->r,
        g->p->color->g,
        g->p->color->b,
        g->p->color->a)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        g->error = 1;
    }

    SDL_RenderFillRect(*g->renderer, g->p->rect);

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
    destroy_player(g->p);
    free(g);

    SDL_Quit();
}
