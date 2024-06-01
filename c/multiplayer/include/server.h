#pragma once

#include "player.h"

typedef struct server {
    player **players;
    size_t nb_players;
} server;


server *create_server();
void destroy_server(server *s);
int serve(server *s);
