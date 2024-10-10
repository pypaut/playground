CC = g++
CFLAGS = -Wall -Werror -Wextra -pedantic -Iinclude
LIBS = -lSDL2 -lm
OBJS = \
	   src/main.cpp \
	   src/error.cpp
BIN = bin

all: build
	./$(BIN)

build:
	$(CC) $(CFLAGS) $(OBJS) -o $(BIN) $(LIBS)

clean:
	$(RM) $(BIN)
