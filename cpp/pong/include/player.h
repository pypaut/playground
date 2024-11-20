#pragma once

#include <SDL2/SDL.h>

class Player {
  public:
    Player();
    ~Player();

    void BuildDefaultPlayer1(int win_height, int win_width);
    void BuildDefaultPlayer2(int win_height, int win_width);

    void SetColor(Uint8 r, Uint8 g, Uint8 b, Uint8 a);
    void SetRect(int x, int y, int w, int h);
    void SetKeys(Uint8 up_key, Uint8 down_key);
    void SetSpeed(int speed);

    void Update(const Uint8 *keys, int win_height, Uint64 dt);
    int Draw(SDL_Renderer *renderer);

  private:
    SDL_Rect *rect;

    int speed;

    Uint8 color_r;
    Uint8 color_g;
    Uint8 color_b;
    Uint8 color_a;

    Uint8 up_key;
    Uint8 down_key;
};
