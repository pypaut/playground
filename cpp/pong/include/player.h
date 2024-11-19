#pragma once

#include <SDL2/SDL.h>


class Player {
    public:
        Player();
        ~Player();

        void SetColor(Uint8 r, Uint8 g, Uint8 b, Uint8 a);
        void SetRect(int x, int y, int w, int h);
        void SetKeys(Uint8 up_key, Uint8 down_key);

        void Update(const Uint8 *keys, int win_height);
        void MoveUp();
        void MoveDown(int window_height);

        int Draw(SDL_Renderer *renderer);

    private:
        SDL_Rect *rect;

        Uint8 color_r;
        Uint8 color_g;
        Uint8 color_b;
        Uint8 color_a;

        Uint8 up_key;
        Uint8 down_key;
};
