#pragma once

#include <SDL2/SDL.h>

class Clock {
    public:
        Clock();
        ~Clock();
        void Tick();

    private:
        Uint64 current_tick;
};
