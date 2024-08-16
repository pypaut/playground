CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99
HEADERS = headers
MAIN = src/main.c
CFILES = src/list.c
BIN = bin

TEST_MAIN = test/main.c
TEST_CFILES = test/test.c \
			  test/test_free_list.c \
			  test/test_new_list.c \
			  test/test_pop_front.c \
			  test/test_push_front.c
TEST_BIN = test_list

all: build
	./$(BIN)

build:
	$(CC) $(CFLAGS) $(CFILES) $(MAIN) -I$(HEADERS) -o $(BIN)

debug: build_debug
	gdb ./$(BIN)

build_debug:
	$(CC) $(CFLAGS) -I$(HEADERS) -g $(CFILES) $(MAIN) -o $(BIN)

test: build_test
	./${TEST_BIN}

test_debug: build_test
	gdb ./${TEST_BIN}

build_test:
	${CC} ${CFLAGS} ${TEST_MAIN} ${TEST_CFILES} ${CFILES} -I$(HEADERS) -Itest -o ${TEST_BIN} -g

clean:
	$(RM) $(BIN)
