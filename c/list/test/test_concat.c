#include "test.h"

int test_concat() {
    log_title("push_back()");
    struct list *my_list1 = new_list();
    struct list *my_list2 = new_list();

    struct list *result = concat(my_list1, my_list2);
    if (test((int)(!result), "result should be NULL")) {
        return 1;
    }
    free_list(result);

    push_back(&my_list1, 1);
    push_back(&my_list2, 2);
    result = concat(my_list1, my_list2);
    if (test((int)(len(result) == 2), "len should be 2")) {
        return 1;
    }
    free_list(result);

    push_back(&my_list1, 1);
    push_back(&my_list2, 2);
    push_back(&my_list2, 2);
    push_back(&my_list2, 2);
    push_back(&my_list2, 2);
    result = concat(my_list1, my_list2);
    if (test((int)(len(result) == 7), "len should be 7")) {
        return 1;
    }
    free_list(result);

    return 0;
}
