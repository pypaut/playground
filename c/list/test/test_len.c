#include "test.h"

int test_len() {
    log_title("len()");

    struct list *my_list = new_list();
    if (test((int)(len(my_list) == 0), "len should be 0")) {
        return 1;
    }

    my_list = push_front(my_list, 3);
    if (test((int)(len(my_list) == 1), "len should be 1")) {
        return 1;
    }

    my_list = push_front(my_list, 3);
    my_list = push_front(my_list, 3);
    my_list = push_front(my_list, 3);
    my_list = push_front(my_list, 3);
    my_list = push_front(my_list, 3);
    if (test((int)(len(my_list) == 6), "len should be 6")) {
        return 1;
    }

    pop_front(&my_list);
    if (test((int)(len(my_list) == 5), "len should be 5")) {
        return 1;
    }

    return 0;
}
