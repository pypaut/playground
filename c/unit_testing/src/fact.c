#include "fact.h"
#include <stdio.h>

int fact(int n) {
    if (n < 0) {
        fprintf(stderr, "error: fact input was %d < 0\n", n);
        return -1;
    }

    switch (n) {
        case 0:
            return 1;
        case 1:
            return 1;
        default:
            return n * fact(n - 1);
    }

    return -1;
}
