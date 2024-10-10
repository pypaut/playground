#include <error.h>

using namespace std;


int main() {
    int W = 1920;
    int H = 1080;

    SDL_Init(SDL_INIT_VIDEO);

    // Create window
    SDL_Window *window = SDL_CreateWindow(
        "Pong", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, W, H, 0
    );
    if (!window) {
        log_error("SDL_CreateWindow");
        return 1;
    }

    // Create renderer
    SDL_Renderer *renderer = SDL_CreateRenderer(
        window, -1, 0
    );
    if (!renderer) {
        log_error("SDL_CreateRenderer");
        return 1;
    }

    // Main loop
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

        // Draw window
        if (SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255)) {
            log_error("SDL_SetRenderDrawColor");
            break;
        }

        if (SDL_RenderClear(renderer)) {
            log_error("SDL_RenderClear");
            break;
        }

        SDL_RenderPresent(renderer);
    }


    SDL_Quit();
    return 0;
}
