#include "fact.h"
#include <stdio.h>


int main() {
    int arg = 10;
    int result = fact(arg);

    printf("fact(%d) is %d\n", arg, result);

    return 0;
}
