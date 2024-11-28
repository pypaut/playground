CC = g++
CXXFLAGS = -Wall -Werror -Wextra -pedantic -Iinclude -std=c++26 -O3
CXX_FLAGS += -m64 -march=native
CXX_FLAGS += -fopt-info-vec-optimized #-fopt-info-vec-missed -ftree-vectorize
ASAN_FLAGS = -fsanitize=address -fsanitize=leak
VALGRIND_FLAGS = --leak-check=full --show-leak-kinds=all --track-origins=yes
# LIBS = -lSDL2 -lm
LIBS = -lGL -lGLEW -lglut -lpthread -lglfw
OBJS = src/main.cpp
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
	$(CC) $(CXXFLAGS) $(OBJS) -o $(BIN) $(LIBS)

build_debug:
	$(CC) $(CXXFLAGS) -g $(OBJS) -o $(BIN) $(LIBS)

build_leaks:
	$(CC) $(CXXFLAGS) -g $(ASAN_FLAGS) $(OBJS) -o $(BIN) $(LIBS)

fmt:
	sh format.sh

clean:
	$(RM) $(BIN)
