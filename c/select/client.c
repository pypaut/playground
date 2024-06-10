#include <stdio.h> 
#include <string.h>
#include <stdlib.h> 
#include <errno.h> 
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/types.h> 
#include <sys/socket.h> 
#include <netinet/in.h> 
#include <sys/time.h> 

#define PORT 8888

int main() {
    /* Variables */
    char *buffer = calloc(1024, sizeof(char));
    char *dir = "Message from the client\0";

    /* Socket components */
    int client_socket_fd;
    struct sockaddr_in serv_addr;

    // Socket creation
    if ((client_socket_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
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
    int status = connect(client_socket_fd, (struct sockaddr*)&serv_addr, sizeof(serv_addr));
    if (status < 0) {
        perror("connect");
        exit(EXIT_FAILURE);
    }

    /* Main loop */
    for (size_t i = 0; i < 4; i++) {
        sleep(1);

        /* Receive from server */
        memset(buffer, 0, 1024);
        read(client_socket_fd, buffer, 1023);
        printf("Received from server: \"%s\"\n", buffer);

        /* Send to server */
        send(client_socket_fd, dir, strlen(dir), 0);
    }

    free(buffer);
    close(client_socket_fd);
    return 0;
}
