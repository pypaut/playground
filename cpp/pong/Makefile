CC = g++
CFLAGS = -Wall -Werror -Wextra -pedantic -Iinclude
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

build:
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

build_debug:
	$(CC) $(CFLAGS) -g $(OBJS) -o $(BIN) $(LIBS)

fmt:
	sh format.sh

check_leaks:
	valgrind --leak-check=full --show-leak-kinds=all --track-origins=yes \
		--verbose --log-file=memcheck.log ./$(BIN)

clean:
	$(RM) $(BIN)
