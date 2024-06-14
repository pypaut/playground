#pragma once

#include <stddef.h>
#include <stdio.h>
#include <string.h>

#include "constants.h"
#include "extract.h"
#include "player_pos.h"

void build_server_message(char *message, player_pos **positions, size_t player_i);
void parse_server_message(player_pos **positions, char *message);
