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

    if (test((int)(first_value == 42), "value should be 42")) {
        return 1;
    }

    if (test((int)(second_value == 213), "value should be 213")) {
        return 1;
    }

    if (test((int)(third_value == 1), "value should be 1")) {
        return 1;
    }

    return 0;
}
