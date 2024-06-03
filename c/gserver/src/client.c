#include "client.h"

int main()
{
    char buffer[1024] = { 0 };
    char* dir = "x:1,y:0";

    Uint64 current_tick = 0;

    int client_socket_fd;
    init_client(&client_socket_fd);

    for (;;) {
        current_tick = clock_tick(current_tick);

        // Send direction/input to server
        send(client_socket_fd, dir, strlen(dir), 0);
        read(client_socket_fd, buffer, 1024 - 1);
        printf("%s\n", buffer);
    }

    // Close the connected socket
    close(client_socket_fd);

    return 0;
}

void init_client(int *client_socket_fd) {
    struct sockaddr_in serv_addr;

    // Socket creation
    if ((*client_socket_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("socket creation");
        exit(EXIT_FAILURE);
    }

    // Setup server address
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(PORT);

    // Convert IPv4 and IPv6 addresses to binary form
    if (inet_pton(AF_INET, "127.0.0.1", &serv_addr.sin_addr) <= 0) {
        perror("invalid address/not supported");
        exit(EXIT_FAILURE);
    }

    // Connect to server
    int status = connect(*client_socket_fd, (struct sockaddr*)&serv_addr, sizeof(serv_addr));
    if (status < 0) {
        perror("connect");
        exit(EXIT_FAILURE);
    }
}

Uint64 clock_tick(Uint64 current_tick) {
    Uint64 last_tick = current_tick;
    current_tick = SDL_GetTicks();
    Uint64 dt = current_tick - last_tick;
    if (dt < 1000 / 60) {
        SDL_Delay(1000 / 60 - dt);
    }

    return current_tick;
}
