#include "test.h"

int test_free_list() {
    log_title("free_list()");
    int ret_value = 0;

    struct list *my_list = new_list();
    free_list(my_list);
    log_success("should not throw error lol");

    return ret_value;
}
