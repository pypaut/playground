#include "message.h"


/* Server side */
void build_server_message(char *message, player_pos **positions) {
    /* Current player in first position in the string */
    char *tmp = calloc(1024, sizeof(char));

    /* All the other players */
    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        if (!positions[i] || !positions[i]->enabled) {
            continue;
        }

        sprintf(tmp, "x:%f,y:%f;", positions[i]->x, positions[i]->y);
        strcat(message, tmp);
    }

    free(tmp);
}

/* Client side */
void parse_server_message(player_pos **positions, char *message) {
    if (!strlen(message)) {
        return;
    }

    char *tok = strtok(message, ";");
    size_t i = 0;

    while (tok) {
        positions[i]->enabled = 1;
        extract_x_y(tok, &(positions[i]->x), &(positions[i]->y));
        tok = strtok(NULL, ";");
        i++;
    }
}
