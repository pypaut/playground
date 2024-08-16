#include "test.h"

int main() {
    if (test_new_list()) {
        return 1;
    }

    if (test_free_list()) {
        return 1;
    }

    return 0;
}
