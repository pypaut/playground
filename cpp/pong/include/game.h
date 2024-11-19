#pragma once

#include <SDL2/SDL.h>
#include <clock.h>
#include <player.h>


class Game {
    public:
        Game(int W, int H);
        ~Game();
        int Draw();
        int Run();

    private:
        int W;
        int H;

        SDL_Window *window;
        SDL_Renderer *renderer;

        Clock *clock;
        Player *player1;
};
