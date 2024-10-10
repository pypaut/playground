#include <SDL2/SDL.h>
#include <error.h>
#include <game.h>
#include <player.h>


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

    // Player 1
    this->player1 = new Player();
    this->player1->SetColor(255, 255, 255, 255);
    this->player1->SetRect(200, 200, 200, 200);
}

Game::~Game() {
    delete(this->clock);
    SDL_Quit();
}

int Game::Draw() {
    if (SDL_SetRenderDrawColor(this->renderer, 0, 0, 0, 255)) {
        log_error("Game::Draw::SDL_SetRenderDrawColor");
        return 1;
    }

    if (SDL_RenderClear(this->renderer)) {
        log_error("Game::Draw::SDL_RenderClear");
        return 1;
    }

    if (this->player1->Draw(this->renderer)) {
        return 1;
    }

    SDL_RenderPresent(renderer);
    return 0;
}

int Game::Run() {
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

        if (this->Draw()) {
            return 1;
        }
    }

    return 0;
}
