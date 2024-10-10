#include <SDL2/SDL.h>
#include <error.h>
#include <game.h>


Game::Game(int W, int H) {
    SDL_Init(SDL_INIT_VIDEO);

    // Window
    this->window = SDL_CreateWindow(
        "Pong", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, W, H, 0
    );
    if (!this->window) {
        log_error("SDL_CreateWindow");
    }

    // Renderer
    this->renderer = SDL_CreateRenderer(window, -1, 0);
    if (!this->renderer) {
        log_error("SDL_CreateRenderer");
    }

    // Clock
    this->clock = new Clock();
}

Game::~Game() {
    delete(this->clock);
    SDL_Quit();
}

void Game::Draw() {
    if (SDL_SetRenderDrawColor(this->renderer, 0, 0, 0, 255)) {
        log_error("SDL_SetRenderDrawColor");
    }

    if (SDL_RenderClear(this->renderer)) {
        log_error("SDL_RenderClear");
    }

    SDL_RenderPresent(renderer);
}

void Game::Run() {
    while (true) {
        // Check for quit event
        SDL_Event event;
        bool should_quit = false;
        while (SDL_PollEvent(&event)) {
            if (event.type == SDL_QUIT) {
                should_quit = true;
                break;
            }
        }
        if (should_quit) {
            break;
        }

        this->Draw();
    }
}
