#include "include.h"


FILE *load_input(char *input_file) {
    FILE *fp = fopen(input_file, "r");
    if (fp == NULL) {
      perror("Error while opening the file.\n");
      exit(EXIT_FAILURE);
    }

    return fp;
}
