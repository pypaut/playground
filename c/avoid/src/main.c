#include <game.h>
#include <clock.h>


int main() {
    int W = 1000;
    int H = 800;

    /* FIXME This isn't working. Goal : hide mouse.
    SDL_Cursor *cursor;
    int32_t cursorData[2] = {0, 0};
    cursor = SDL_CreateCursor((Uint8 *)cursorData, (Uint8 *)cursorData, 8, 8, 4, 4);
    SDL_SetCursor(cursor);
    SDL_ShowCursor(SDL_DISABLE);
    */

    game *g = init_game(W, H);

    while (g->is_running && !g->error) {
        handle_quit_event(g);
        update_game(g);
        draw_game(g);
    }

    destroy_game(g);
    return 0;
}
