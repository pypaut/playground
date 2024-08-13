CC = g++
CFLAGS = -Wall -Werror -Wextra -pedantic -Iinclude
OBJS = main.cpp
BIN = bin

all: build
	./$(BIN)

debug: build_debug
	gdb ./$(BIN)

build:
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

build_debug:
	$(CC) $(CFLAGS) -g $(OBJS) -o $(BIN) $(LIBS)

clean:
	$(RM) $(BIN)
