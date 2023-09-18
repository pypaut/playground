#include "fact.h"
#include <assert.h>
#include <stdio.h>

#define ANSI_COLOR_GREEN   "\x1b[32m"
#define ANSI_COLOR_RESET   "\x1b[0m"

int main() {
    assert(fact(0) == 1);
    assert(fact(1) == 1);
    assert(fact(2) == 2);
    assert(fact(3) == 6);
    assert(fact(4) == 24);
    printf(ANSI_COLOR_GREEN   "OK!"   ANSI_COLOR_RESET "\n");
    return 0;
}
