#pragma once

#include <stdio.h>
#include "list.h"

#define ANSI_COLOR_GREEN   "\x1b[32m"
#define ANSI_COLOR_RED     "\x1b[31m"
#define ANSI_COLOR_RESET   "\x1b[0m"

void log_title(char *msg);
void log_err(char *msg);
void log_success(char *msg);

int test(int condition, char *message);

int test_free_list();
int test_new_list();

int test_at();
int test_len();

int test_push_front();
int test_pop_front();
int test_pop_at();

int test_pprint();
