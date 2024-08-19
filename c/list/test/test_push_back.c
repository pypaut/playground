#include "test.h"

int test_push_back() {
    log_title("push_back()");

    struct list *my_list = new_list();

    push_back(&my_list, 1);
    if (test((int)(my_list->value == 1), "value should be 1")) {
        return 1;
    }
    if (test((int)(len(my_list) == 1), "len should be 1")) {
        return 1;
    }

    push_back(&my_list, -1);
    if (test((int)(my_list->value == 1), "value should be 1")) {
        return 1;
    }
    if (test((int)(len(my_list) == 2), "len should be 2")) {
        return 1;
    }

    push_back(&my_list, 83429);
    if (test((int)(my_list->value == 1), "value should be 1")) {
        return 1;
    }
    if (test((int)(len(my_list) == 3), "len should be 3")) {
        return 1;
    }

    if (test((int)(at(my_list, 0) == 1), "value should be 1")) {
        return 1;
    }

    if (test((int)(at(my_list, 1) == -1), "value should be 1")) {
        return 1;
    }

    if (test((int)(at(my_list, 2) == 83429), "value should be 83429")) {
        return 1;
    }

    return 0;
}
