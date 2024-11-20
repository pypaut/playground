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

    this->player1 = new Player();
    this->player2 = new Player();

    this->player1->BuildDefaultPlayer1(H, W);
    this->player2->BuildDefaultPlayer2(H, W);

    this->ball = new Ball();
    this->ball->Build(W, H);
}

Game::~Game() {
    std::cout << "Freeing game resources" << std::endl;
    delete(this->clock);
    delete(this->player1);
    delete(this->player2);
    delete(this->ball);
    SDL_Quit();
}

void Game::Update(const Uint8 *keys) {
    Uint64 dt = this->clock->Tick();
    this->ball->Update(this->W, this->H, dt);
    this->player1->Update(keys, this->H, dt);
    this->player2->Update(keys, this->H, dt);
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

    if (this->ball->Draw(this->renderer)) {
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
