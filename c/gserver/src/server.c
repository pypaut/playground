#include "server.h"


int main() {
    char *buffer = calloc(1024, sizeof(char));
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
        /* Receive from client */
        memset(buffer, 0, 1024);
        read(client_socket_fd, buffer, 1023);

        /* Extract dir */
        float dir_x = 0;
        float dir_y = 0;
        extract_dir(buffer, &dir_x, &dir_y);

        /* Update client position */

        /* Send to client */
        sprintf(pos, "x:%d,y:%d", pos_x, pos_y);
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

void extract_dir(char *buffer, float *dir_x, float *dir_y) {
    /* Format : "x:0.000000,y:0.0000000" */

    /* Extract direction from string, as string */
    char *dir_x_string = calloc(20, sizeof(char));
    char *dir_y_string = calloc(20, sizeof(char));
    extract_dir_str(buffer, dir_x_string, dir_y_string);

    /* Convert to float */
    *dir_x = atof(dir_x_string);
    *dir_y = atof(dir_y_string);
    printf("%f,%f\n", *dir_x, *dir_y);

    free(dir_x_string);
    free(dir_y_string);
}

void extract_dir_str(char *buffer, char *dir_x_string, char *dir_y_string) {
    size_t buffer_i = 0;
    while (buffer[buffer_i] != ':') {
        buffer_i++;
    }
    buffer_i++;

    size_t dir_i = 0;
    while (buffer[buffer_i] != ',') {
        dir_x_string[dir_i] = buffer[buffer_i];
        buffer_i++;
        dir_i++;
    }
    buffer_i++;

    while (buffer[buffer_i] != ':') {
        buffer_i++;
    }
    buffer_i++;

    dir_i = 0;
    while (buffer[buffer_i]) {
        dir_y_string[dir_i] = buffer[buffer_i];
        buffer_i++;
        dir_i++;
    }
}
