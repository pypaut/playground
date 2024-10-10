#pragma once

#include <SDL2/SDL.h>


class Game {
    public:
        Game(int W, int H);
        ~Game();
        void Draw();
        void Run();

    private:
        SDL_Window *window;
        SDL_Renderer *renderer;
};
