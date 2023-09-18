CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99
INCLUDE = include
MAIN = src/main.c
TEST_MAIN = test/test_fact.c
CFILES = src/fact.c
BIN = fact
TEST_BIN = test_fact

all: build
	./${BIN}

build:
	${CC} ${CFLAGS} ${MAIN} ${CFILES} -I${INCLUDE} -o ${BIN}

test: build_test
	./${TEST_BIN}

build_test:
	${CC} ${CFLAGS} ${TEST_MAIN} ${CFILES} -I${INCLUDE} -o ${TEST_BIN} -g

clean:
	${RM} ${BIN} ${TEST_BIN}
