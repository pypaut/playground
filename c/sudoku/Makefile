CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99 -Iinclude -fsanitize=address
LIBS = -lasan
OBJS = src/main.c \
	   src/grid.c
INCLUDE = include
BIN = sudoku

all: build
	./$(BIN)

build:
	$(CC) $(CFLAGS) $(OBJS) -I$(INCLUDE) -o $(BIN) $(LIBS)

clean:
	$(RM) $(BIN)
