#include "test.h"

int test_push_front() {
    log_title("push_front()");

    struct list *my_list = new_list();
    my_list = push_front(my_list, 1);

    if (my_list->value != 1) {
        log_err("value should be 1");
        return 1;
    } else {
        log_success("value should be 1");
    }

    my_list = push_front(my_list, -1);

    if (my_list->value != -1) {
        log_err("value should be -1");
        return 1;
    } else {
        log_success("value should be -1");
    }

    my_list = push_front(my_list, 83429);

    if (my_list->value != 83429) {
        log_err("value should be 83429");
        return 1;
    } else {
        log_success("value should be 83429");
    }

    return 0;
}
