#include "test.h"

int test_push_front() {
    log_title("push_front()");

    struct list *my_list = new_list();

    my_list = push_front(my_list, 1);
    if (test((int)(my_list->value == 1), "value should be 1")) {
        return 1;
    }

    my_list = push_front(my_list, -1);
    if (test((int)(my_list->value == -1), "value should be -1")) {
        return 1;
    }

    my_list = push_front(my_list, 83429);
    if (test((int)(my_list->value == 83429), "value should be 83429")) {
        return 1;
    }

    return 0;
}
