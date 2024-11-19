#include <SDL2/SDL.h>
#include <error.h>
#include <game.h>
#include <player.h>


Game::Game(int W, int H) {
    SDL_Init(SDL_INIT_VIDEO);

    // Window
    this->W = W;
    this->H = H;

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

    int player_w = W / 100;
    int player_x = W / 20;

    int player_h = H / 5;
    int player_y = (H - player_h) / 2;

    this->player1->SetRect(player_x, player_y, player_w, player_h);

    // Player 2
    this->player2 = new Player();
    this->player2->SetColor(255, 255, 255, 255);

    this->player2->SetRect(
            W - player_x - player_w,
            player_y,
            player_w,
            player_h
    );
}

Game::~Game() {
    delete(this->clock);
    SDL_Quit();
}

void Game::Update(const Uint8 *keys) {
    this->player1->Update(keys, this->H);
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

    if (this->player2->Draw(this->renderer)) {
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

        // Input
        const Uint8 *keys = SDL_GetKeyboardState(NULL);
        this->Update(keys);

        // Draw
        if (this->Draw()) {
          return 1;
        }
    }

    return 0;
}
