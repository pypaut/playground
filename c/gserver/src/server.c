#include "server.h"


int main() {
    char *buffer = calloc(1024, sizeof(char));
    float pos_x = 0;
    float pos_y = 0;
    char *pos = calloc(256, sizeof(char));

    int server_socket_fd;
    struct sockaddr_in address;
    socklen_t addrlen = sizeof(address);
    init_server(&server_socket_fd, &address);

    int client_socket_fd;
    if ((client_socket_fd = accept(server_socket_fd, (struct sockaddr*)&address, &addrlen)) < 0) {
        perror("accept");
        exit(EXIT_FAILURE);
    }

    for (;;) {
        /* Receive from client */
        memset(buffer, 0, 1024);
        read(client_socket_fd, buffer, 1023);

        /* Extract dir */
        float dir_x = 0;
        float dir_y = 0;
        extract_x_y(buffer, &dir_x, &dir_y);

        /* Update client position */
        pos_x += dir_x * PLAYER_SPEED * SCALE;
        pos_y += dir_y * PLAYER_SPEED * SCALE;
        pos_x = clamp(pos_x, 0, (W - PLAYER_SIZE) * SCALE);
        pos_y = clamp(pos_y, 0, (H - PLAYER_SIZE) * SCALE);

        /* Send to client */
        sprintf(pos, "x:%f,y:%f", pos_x, pos_y);
        send(client_socket_fd, pos, strlen(pos), 0);
    }
 
    free(pos);
    free(buffer);
    close(client_socket_fd);
    close(server_socket_fd);

    return 0;
}

void init_server(int *server_socket_fd, struct sockaddr_in *address) {
    int opt = 1;
 
    // Create socket
    if ((*server_socket_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("socket failed");
        exit(EXIT_FAILURE);
    }
 
    // Set socket options
    if (setsockopt(*server_socket_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt))) {
        perror("setsockopt");
        exit(EXIT_FAILURE);
    }

    // Configure address
    address->sin_family = AF_INET;
    address->sin_addr.s_addr = INADDR_ANY;
    address->sin_port = htons(PORT);
 
    if (bind(*server_socket_fd, (struct sockaddr*)address, sizeof(*address)) < 0) {
        perror("bind failed");
        exit(EXIT_FAILURE);
    }

    if (listen(*server_socket_fd, 3) < 0) {
        perror("listen");
        exit(EXIT_FAILURE);
    }
}

double clamp(double d, double min, double max) {
  const double t = d < min ? min : d;
  return t > max ? max : t;
}
