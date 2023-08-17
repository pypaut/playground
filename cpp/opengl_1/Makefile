CC = g++

CPP_FILES = src/main.cpp

CXX_FLAGS += -Wall -Wextra -O3 -g -std=c++11
CXX_FLAGS += -m64 -march=native
CXX_FLAGS += -fopt-info-vec-optimized #-fopt-info-vec-missed -ftree-vectorize
LDXX_FLAGS = -lGL -lGLEW -lglut -lpthread -lglfw

BIN = opengl

all: build
	./$(BIN)

build:
	$(CC) $(CPP_FILES) -o $(BIN) $(CXX_FLAGS) $(LDXX_FLAGS)

clean:
	rm -f $(BIN)
