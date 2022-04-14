CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99 -Iinclude
LIBS = -lSDL2 -lm
OBJS = src/main.c \
	   src/player.c \
	   src/ball.c
BIN = bin

all: build
	./$(BIN)

build:
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

clean:
	$(RM) $(BIN)
