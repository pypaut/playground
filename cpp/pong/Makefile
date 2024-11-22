CC = g++
CFLAGS = -Wall -Werror -Wextra -pedantic -Iinclude
ASAN_FLAGS = -fsanitize=address -fsanitize=leak
VALGRIND_FLAGS = --leak-check=full --show-leak-kinds=all --track-origins=yes
LIBS = -lSDL2 -lm
OBJS = \
	   src/ball.cpp \
	   src/clock.cpp \
	   src/error.cpp \
	   src/game.cpp \
	   src/main.cpp \
	   src/player.cpp
BIN = bin

all: build
	./$(BIN)

debug: build_debug
	gdb ./$(BIN)

leaks: build_leaks
	./$(BIN)

valgrind: build
	valgrind $(VALGRIND_FLAGS) \
		--log-file=memcheck.log ./$(BIN)

build:
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

build_debug:
	$(CC) $(CFLAGS) -g $(OBJS) -o $(BIN) $(LIBS)

build_leaks:
	$(CC) $(CFLAGS) -g $(ASAN_FLAGS) $(OBJS) -o $(BIN) $(LIBS)

fmt:
	sh format.sh

clean:
	$(RM) $(BIN)
