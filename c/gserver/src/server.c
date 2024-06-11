#include "server.h"


int main() {
    char *buffer = calloc(1024, sizeof(char));
    player_pos *positions = calloc(MAX_CLIENTS + 1, sizeof(player_pos));

    int max_sd, activity, new_client_socket;

    // Server
    int server_socket_fd;
    struct sockaddr_in address;
    init_server(&server_socket_fd, &address);

    // Clients
    fd_set readfds;
    int client_sockets_fds[MAX_CLIENTS];
    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        client_sockets_fds[i] = 0;
    }

    for (;;) {
        // Clear the socket set
        FD_ZERO(&readfds);

        // Add master socket to set
        FD_SET(server_socket_fd, &readfds);
        max_sd = server_socket_fd;

        // Add child sockets to set
        add_child_sockets_to_set(&readfds, client_sockets_fds, &max_sd);

        // Wait for an activity on one of the sockets, timeout is NULL,
        // so wait indefinitely
        activity = select(max_sd + 1, &readfds, NULL, NULL, NULL);
        if ((activity < 0) && (errno != EINTR)) {
            fprintf(stderr, "select error");
        }

        if (FD_ISSET(server_socket_fd, &readfds)) {
            handle_new_client(
                &new_client_socket,
                &server_socket_fd,
                client_sockets_fds,
                &address
            );
        } else {
            handle_clients_messages(
                client_sockets_fds,
                &readfds,
                &address,
                positions
            );
        }
    }
 
    free(positions);
    free(buffer);
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
    if (setsockopt(
            *server_socket_fd,
            SOL_SOCKET,
            SO_REUSEADDR,
            &opt,
            sizeof(opt))
    ) {
        perror("setsockopt");
        exit(EXIT_FAILURE);
    }

    // Configure address
    address->sin_family = AF_INET;
    address->sin_addr.s_addr = INADDR_ANY;
    address->sin_port = htons(PORT);
 
    if (
        bind(*server_socket_fd, (struct sockaddr*)address, sizeof(*address)) < 0
    ) {
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

void update_pos(float *pos_x, float *pos_y, float dir_x, float dir_y) {
    *pos_x += dir_x * PLAYER_SPEED * SCALE;
    *pos_y += dir_y * PLAYER_SPEED * SCALE;
    *pos_x = clamp(*pos_x, 0, (W - PLAYER_SIZE) * SCALE);
    *pos_y = clamp(*pos_y, 0, (H - PLAYER_SIZE) * SCALE);
}

void add_child_sockets_to_set(
    fd_set *readfds,
    int *client_sockets,
    int *max_sd
) {
    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        // Socket descriptor
        int sd = client_sockets[i];

        // If valid socket descriptor then add to read list
        if (sd > 0)
            FD_SET(sd, readfds);

        // Highest file descriptor number, need it for the select function
        if (sd > *max_sd)
            *max_sd = sd;
    }
}

void handle_new_client(
    int *new_socket,
    int *server_socket_fd,
    int *client_sockets,
    struct sockaddr_in *address
) {
    int addrlen = sizeof(*address);
    *new_socket = accept(
        *server_socket_fd,
        (struct sockaddr*)address,
        (socklen_t*)&addrlen
    );
    if (*new_socket < 0) {
        perror("accept");
        exit(EXIT_FAILURE);
    }

    // Add new socket to array of sockets
    for (int i = 0; i < MAX_CLIENTS; i++) {
        if (client_sockets[i] == 0) {
            client_sockets[i] = *new_socket;
            printf(
                "[ID: %d, IP: %s, PORT: %d]: connected\n",
                i, inet_ntoa(address->sin_addr), ntohs(address->sin_port)
            );
            break;
        }
    }
}

void handle_clients_messages(
    int *client_sockets,
    fd_set *readfds,
    struct sockaddr_in *address,
    player_pos *positions
) {
    int addrlen = sizeof(address);
    char *buffer = calloc(1024, sizeof(char));

    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        int sd = client_sockets[i];

        if (FD_ISSET(sd, readfds)) {
            // If empty message, closing connection
            if (read(sd, buffer, 1024) == 0) {
                // Somebody disconnected , get his details and print
                getpeername(
                    sd,
                    (struct sockaddr*)address,
                    (socklen_t*)&addrlen
                );
                printf(
                    "[ID: %ld, IP: %s, PORT: %d]: disconnected\n",
                    i, inet_ntoa(address->sin_addr), ntohs(address->sin_port)
                );

                // Close the socket and mark as 0 in list for reuse
                close(sd);
                client_sockets[i] = 0;
            } else {
                // The buffer contains data from client
                printf(
                    "[ID: %ld, IP: %s, PORT: %d]: read \"%s\"\n",
                    i,
                    inet_ntoa(address->sin_addr),
                    ntohs(address->sin_port),
                    buffer
                );

                /* Extract dir */
                float dir_x = 0;
                float dir_y = 0;
                extract_x_y(buffer, &dir_x, &dir_y);

                /* Update client position */
                update_pos(
                    &(positions[i].pos_x),
                    &(positions[i].pos_y),
                    dir_x,
                    dir_y
                );

                /* Send to client */
                char *pos = calloc(1024, sizeof(char));
                sprintf(pos, "x:%f,y:%f", positions[i].pos_x, positions[i].pos_y);
                send(sd, pos, strlen(pos), 0);
                printf(
                    "[ID: %ld, IP: %s, PORT: %d]: send \"%s\"\n",
                    i,
                    inet_ntoa(address->sin_addr),
                    ntohs(address->sin_port),
                    pos
                );
                free(pos);
            }
        }
    }

    free(buffer);
}
