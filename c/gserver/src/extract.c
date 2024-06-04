#include "extract.h"


void extract_x_y(char *buffer, float *dir_x, float *dir_y) {
    /* Format : "x:0.000000,y:0.0000000" */

    /* Extract direction from string, as string */
    char *dir_x_string = calloc(20, sizeof(char));
    char *dir_y_string = calloc(20, sizeof(char));
    extract_x_y_str(buffer, dir_x_string, dir_y_string);

    /* Convert to float */
    *dir_x = atof(dir_x_string);
    *dir_y = atof(dir_y_string);
    printf("%f,%f\n", *dir_x, *dir_y);

    free(dir_x_string);
    free(dir_y_string);
}

void extract_x_y_str(char *buffer, char *dir_x_string, char *dir_y_string) {
    size_t buffer_i = 0;
    while (buffer[buffer_i] != ':') {
        buffer_i++;
    }
    buffer_i++;

    size_t dir_i = 0;
    while (buffer[buffer_i] != ',') {
        dir_x_string[dir_i] = buffer[buffer_i];
        buffer_i++;
        dir_i++;
    }
    buffer_i++;

    while (buffer[buffer_i] != ':') {
        buffer_i++;
    }
    buffer_i++;

    dir_i = 0;
    while (buffer[buffer_i]) {
        dir_y_string[dir_i] = buffer[buffer_i];
        buffer_i++;
        dir_i++;
    }
}
