#include <clock.h>

Uint64 tick(clock *c) {
    Uint64 last_tick = c->current_tick;
    c->current_tick = SDL_GetTicks();
    Uint64 dt = c->current_tick - last_tick;
    if (dt < 1000 / 60) {
        SDL_Delay(1000 / 60 - dt);
    }

    return dt;
}

clock *new_clock() {
    clock *c = calloc(1, sizeof(clock));
    return c;
}

void destroy_clock(clock *c) {
    free(c);
}
