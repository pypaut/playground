CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99 -Iinclude
LIBS = -lSDL2 -lm
OBJS = src/main.c \
	   src/game.c \
	   src/clock.c \
	   src/player.c
BIN = game
SERVER_OBJS = src/server.c
SERVER_BIN = server

all: build
	./$(BIN)

build:
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

build_server:
	$(CC) $(CFLAGS) $(SERVER_OBJS) -o $(SERVER_BIN)

clean:
	$(RM) $(BIN)
