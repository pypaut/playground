CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99 -Iinclude
LIBS = -lSDL2 -lm

CLIENT_OBJS = src/client.c
CLIENT_BIN = client
SERVER_OBJS = src/server.c
SERVER_BIN = server

all: build_client build_server

build_client:
	$(CC) $(CFLAGS) $(CLIENT_OBJS) -o $(CLIENT_BIN) $(LIBS)

build_server:
	$(CC) $(CFLAGS) $(SERVER_OBJS) -o $(SERVER_BIN) $(LIBS)

clean:
	$(RM) $(BIN)
