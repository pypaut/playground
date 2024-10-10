#pragma once

#include <SDL2/SDL.h>


class Player {
    public:
        Player();
        ~Player();
        void SetColor(Uint8 r, Uint8 g, Uint8 b, Uint8 a);
        void SetRect(int x, int y, int w, int h);
        int Draw(SDL_Renderer *renderer);

    private:
        SDL_Rect *rect;
        Uint8 color_r;
        Uint8 color_g;
        Uint8 color_b;
        Uint8 color_a;
};
