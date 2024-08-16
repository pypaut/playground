#include "test.h"

int test_at() {
    log_title("at()");

    struct list *my_list = new_list();
    my_list = push_front(my_list, 1);
    my_list = push_front(my_list, 213);
    my_list = push_front(my_list, 42);

    if (test((int)(at(my_list, 0) == 42), "value should be 42")) {
        return 1;
    }

    if (test((int)(at(my_list, 1) == 213), "value should be 213")) {
        return 1;
    }

    if (test((int)(at(my_list, 2) == 1), "value should be 1")) {
        return 1;
    }

    return 0;
}
