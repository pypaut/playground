#pragma once

#include <SDL2/SDL.h>

class Clock {
  public:
    Clock();
    ~Clock();
    Uint64 Tick();

  private:
    Uint64 current_tick;
};
