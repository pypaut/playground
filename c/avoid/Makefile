CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99 -Iinclude
LIBS = -lSDL2 -lm
OBJS = src/main.c \
	   src/clock.c \
	   src/game.c \
	   src/player.c
BIN = bin

all: build
	./$(BIN)

build:
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

clean:
	$(RM) $(BIN)
