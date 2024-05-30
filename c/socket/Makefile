CC = gcc
CFLAGS = -Wall -Werror -Wextra -pedantic -std=c99

all: build_server build_client
	echo "Build all"

build_client:
	$(CC) $(CFLAGS) client.c -o client

build_server:
	$(CC) $(CFLAGS) server.c -o server
