#include "include.h"


int main() {
    FILE *fp = load_input("test.txt");
    int score = compute_score(fp);
    printf("Score is %d\n", score);
    return 0;
}
