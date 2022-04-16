#pragma once

#include <SDL2/SDL.h>

typedef struct clock {
    Uint64 current_tick;
} clock;

void tick(clock *c);
clock *new_clock();
void destroy_clock(clock *c);
