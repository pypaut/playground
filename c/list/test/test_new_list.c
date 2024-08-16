#include "test.h"

int test_new_list() {
    log_title("new_list()");

    struct list *my_list = new_list();

    if (test((int)(my_list == NULL), "my_list should be NULL")) {
        return 1;
    }

    return 0;
}
