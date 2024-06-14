#include "message.h"


/* Server side */
void build_server_message(char *message, player_pos **positions, size_t player_i) {
    /* Current player in first position in the string */
    // sprintf(message, "x:%f,y:%f;", positions[player_i].x, positions[player_i].y); NOT NECESSARY
    player_i = player_i;

    char *tmp = calloc(1024, sizeof(char));

    /* All the other players */
    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        // if (!positions[i] || i == player_i) {
        if (!positions[i]) {
            continue;
        }

        sprintf(tmp, "x:%f,y:%f;", positions[i]->x, positions[i]->y);
        // sprintf(message, "%sx:%f,y:%f;", message, positions[i]->x, positions[i]->y);
        strcat(tmp, message);
    }

    free(tmp);
}

/* Client side */
void parse_server_message(player_pos **positions, char *message) {
    char *tok = strtok(message, ";");
    size_t i = 0;

    while (tok) {
        extract_x_y(tok, &(positions[i]->x), &(positions[i]->y));
        tok = strtok(NULL, ";");
        i++;
    }
}
