#include "player_pos.h"

player_pos **new_player_pos_list() {
    player_pos **positions = calloc(MAX_CLIENTS, 8);
    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        positions[i] = calloc(1, sizeof(player_pos));
    }

    return positions;
}

void free_player_pos_list(player_pos **positions) {
    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        free(positions[i]);
    }

    free(positions);
}
