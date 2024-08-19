#include "test.h"

int test_pprint() {
    log_title("pprint()");

    struct list *l = new_list();
    push_front(&l, 2);
    push_front(&l, 2);
    push_front(&l, 2);
    push_front(&l, 3);
    pprint(l);

    return 0;
}
