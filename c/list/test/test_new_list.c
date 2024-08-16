#include "test.h"

int test_new_list() {
    log_title("new_list()");

    struct list *my_list = new_list();

    if (test((int)(my_list != NULL), "my_list should not be NULL")) {
        return 1;
    }

    if (test((int)(my_list->next == NULL), "my_list->next should be NULL")) {
        return 1;
    }

    if (test((int)(my_list->value == 0), "value should be 0")) {
        return 1;
    }

    return 0;
}
