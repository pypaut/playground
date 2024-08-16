#include "test.h"

int test_new_list() {
    log_title("new_list()");

    struct list *my_list = new_list();
    if (!my_list) {
        log_err("list should not be NULL");
        return 1;
    } else {
        log_success("list should not be NULL");
    }

    if (my_list->next) {
        log_err("list->next should be NULL");
        return 1;
    } else {
        log_success("list->next should be NULL");
    }

    if (my_list->value) {
        log_err("list->value should be 0");
        return 1;
    } else {
        log_success("list->value should be 0");
    }

    return 0;
}
