#include "test.h"

int test_pop_front() {
    log_title("pop_front()");

    struct list *my_list = new_list();
    my_list = push_front(my_list, 1);
    my_list = push_front(my_list, 213);
    my_list = push_front(my_list, 42);

    int first_value = pop_front(&my_list);
    int second_value = pop_front(&my_list);
    int third_value = pop_front(&my_list);

    if (first_value != 42) {
        log_err("value should be 42");
        return 1;
    } else {
        log_success("value should be 42");
    }

    if (second_value != 213) {
        log_err("value should be 213");
        return 1;
    } else {
        log_success("value should be 213");
    }

    if (third_value != 1) {
        log_err("value should be 1");
        return 1;
    } else {
        log_success("value should be 1");
    }

    return 0;
}
