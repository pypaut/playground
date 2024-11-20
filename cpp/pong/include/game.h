#pragma once

#include <SDL2/SDL.h>
#include <ball.h>
#include <clock.h>
#include <player.h>

class Game {
  public:
    Game(int W, int H);
    ~Game();

    void Update(const Uint8 *keys);
    int Draw();
    int Run();

  private:
    int W;
    int H;

    SDL_Window *window;
    SDL_Renderer *renderer;

    Clock *clock;
    Player *player1;
    Player *player2;
    Ball *ball;
};
