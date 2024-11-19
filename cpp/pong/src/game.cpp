#include <SDL2/SDL.h>
#include <error.h>
#include <game.h>
#include <player.h>


Player *BuildDefaultPlayer1(int win_height, int win_width) {
    Player *p1 = new Player();
    p1->SetColor(255, 255, 255, 255);

    int player_w = win_width / 100;
    int player_x = win_width / 20;

    int player_h = win_height / 5;
    int player_y = (win_height - player_h) / 2;

    p1->SetRect(player_x, player_y, player_w, player_h);
    p1->SetKeys(SDL_SCANCODE_W, SDL_SCANCODE_S);
    p1->SetSpeed(1);

    return p1;
}

Player *BuildDefaultPlayer2(int win_height, int win_width) {
    Player *p2 = new Player();
    p2->SetColor(255, 255, 255, 255);

    int player_w = win_width / 100;
    int player_x = win_width - (win_width/20) - player_w;

    int player_h = win_height / 5;
    int player_y = (win_height - player_h) / 2;

    p2->SetRect(
            player_x,
            player_y,
            player_w,
            player_h
    );
    p2->SetKeys(SDL_SCANCODE_UP, SDL_SCANCODE_DOWN);
    p2->SetSpeed(1);

    return p2;
}

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

    this->player1 = BuildDefaultPlayer1(H, W);
    this->player2 = BuildDefaultPlayer2(H, W);
}

Game::~Game() {
    std::cout << "Freeing game resources" << std::endl;
    delete(this->clock);
    delete(this->player1);
    delete(this->player2);
    SDL_Quit();
}

void Game::Update(const Uint8 *keys) {
    Uint64 dt = this->clock->Tick();
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
