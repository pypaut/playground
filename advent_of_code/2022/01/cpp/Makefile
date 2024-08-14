CC = g++
CFLAGS = -Wall -Werror -Wextra -pedantic -Iinclude
OBJS = main.cpp
BIN = bin

all: build
	./$(BIN)

debug: build_debug
	gdb ./$(BIN)

build: format
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

build_debug: format
	$(CC) $(CFLAGS) -g $(OBJS) -o $(BIN) $(LIBS)

format:
	clang-format main.cpp > tmp
	mv tmp main.cpp

clean:
	$(RM) $(BIN)
