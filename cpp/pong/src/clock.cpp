#include <SDL2/SDL.h>
#include <clock.h>


Clock::Clock() {
    this->current_tick = 0;
}

Clock::~Clock() {}

void Clock::Tick() {
    Uint64 last_tick = this->current_tick;
    this->current_tick = SDL_GetTicks();

    Uint64 dt = this->current_tick - last_tick;
    if (dt < 1000 / 60) {
        SDL_Delay(1000 / 60 - dt);
    }
}
