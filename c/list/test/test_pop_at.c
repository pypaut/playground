#include "test.h"

int test_pop_at() {
    log_title("pop_at()");

    struct list *my_list = new_list();
    push_front(&my_list, 1);
    push_front(&my_list, 213);
    push_front(&my_list, 42);

    if (test((int)(len(my_list) == 3), "len should be 3")) {
        return 1;
    }

    if (test((int)(pop_at(my_list, 1) == 213), "value should be 213")) {
        return 1;
    }

    if (test((int)(len(my_list) == 2), "len should be 2")) {
        return 1;
    }

    if (test((int)(pop_at(my_list, 1) == 1), "value should be 1")) {
        return 1;
    }

    if (test((int)(len(my_list) == 1), "len should be 1")) {
        return 1;
    }

    if (test((int)(at(my_list, 0) == 42), "value should be 42")) {
        return 1;
    }

    return 0;
}
