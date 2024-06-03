#include "server.h"


int main() {
    char buffer[1024] = { 0 };
    int pos_x = 0;
    int pos_y = 0;
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
        // Get client input
        read(client_socket_fd, buffer, 1024 - 1);
        printf("%s\n", buffer);

        // TODO Update client position

        // Send client position
        sprintf(pos, "x:%d,y:%d", pos_x, pos_y);
        send(client_socket_fd, pos, strlen(pos), 0);
    }
 
    free(pos);
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
