#pragma once

#include <stdlib.h>

#include "constants.h"

typedef struct player_pos {
    float x;
    float y;
} player_pos;

player_pos **new_player_pos_list();
void free_player_pos_list(player_pos **positions);
