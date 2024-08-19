#include "test.h"

int main() {
    if (test_new_list()) {
        return 1;
    }

    if (test_free_list()) {
        return 1;
    }

    if (test_push_front()) {
        return 1;
    }

    if (test_pop_front()) {
        return 1;
    }

    if (test_at()) {
        return 1;
    }

    if (test_len()) {
        return 1;
    }

    if (test_pprint()) {
        return 1;
    }

    return 0;
}
