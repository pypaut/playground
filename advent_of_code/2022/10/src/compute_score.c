#include "include.h"

enum Characters { PAREN, BRACK, BRACE, CMP };

void set_characters_score(int characters_score[]) {
    characters_score[(int)')'] = 3;
    characters_score[(int)']'] = 57;
    characters_score[(int)'}'] = 1197;
    characters_score[(int)'>'] = 25137;
}

char reach_end_of_line(FILE *fp) {
    char c = fgetc(fp);

    while (c != EOF && c != '\n') {
        c = fgetc(fp);
    }

    return c;
}

int compute_score(FILE *fp) {
    int score = 0;

    int characters_score[4096];
    set_characters_score(characters_score);

    char c = fgetc(fp);
    char expected;
    char next_expected;

    while (c != EOF) {
        if (c == '(') {
            next_expected = expected;
            expected = ')';
        }
        else if (c == '[') {
            next_expected = expected;
            expected = ']';
        }
        else if (c == '{') {
            next_expected = expected;
            expected = '}';
        }
        else if (c == '<') {
            next_expected = expected;
            expected = '>';
        }
        else {
            if (c != expected) {
                printf("Corrupted line : unexpected %c\n", c);
                score = score + characters_score[(int)c];
                c = reach_end_of_line(fp);
            } else {
                expected = next_expected;
            }
        }

        if (c != EOF) {
            c = fgetc(fp);
        }

        if (c == '\n') {
            c = fgetc(fp);
        }
    }

    return score;
}
