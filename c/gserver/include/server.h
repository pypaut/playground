#pragma once

#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#include "constants.h"
#include "extract.h"

void init_server(int *server_socket_fd, struct sockaddr_in *address);
double clamp(double d, double min, double max);
