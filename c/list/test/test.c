#include "test.h"

void log_title(char *msg) {
    printf("### TESTS: %s ###\n", msg);
}

void log_err(char *msg) {
    printf(ANSI_COLOR_RED   "%s: KO!"   ANSI_COLOR_RESET "\n", msg);
}

void log_success(char *msg) {
    printf(ANSI_COLOR_GREEN   "%s: OK!"   ANSI_COLOR_RESET "\n", msg);
}
