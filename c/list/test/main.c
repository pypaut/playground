#include "test.h"

int main() {
    int ret_value = 0;

    if (test_new_list()) {
        ret_value = 1;
    }

    if (test_free_list()) {
        ret_value = 1;
    }

    return ret_value;
}
