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
    int opt = 1; 
    int master_socket, addrlen, new_socket, client_socket[30], max_clients=3, activity, i, valread, sd; 
    int max_sd; 
    struct sockaddr_in address; 

    char buffer[1025];

    // Set of socket descriptors 
    fd_set readfds; 

    // Messages
    char *send_buffer = calloc(1024, sizeof(char));

    // Initialise all client_socket[] to 0 so not checked
    for (i = 0; i < max_clients; i++) { 
        client_socket[i] = 0; 
    } 

    // Create a master socket
    if ((master_socket = socket(AF_INET , SOCK_STREAM , 0)) == 0) { 
        perror("socket failed"); 
        exit(EXIT_FAILURE); 
    } 

    // Set master socket to allow multiple connections,
    // this is just a good habit, it will work without this
    if (setsockopt(master_socket, SOL_SOCKET, SO_REUSEADDR, (char *)&opt, sizeof(opt)) < 0) { 
        perror("setsockopt");
        exit(EXIT_FAILURE);
    } 

    // Type of socket created
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(PORT);

    // Bind the socket to localhost port 8888
    if (bind(master_socket, (struct sockaddr *)&address, sizeof(address)) < 0) {
        perror("bind failed");
        exit(EXIT_FAILURE);
    }
    printf("Listener on port %d\n", PORT);

    // Try to specify maximum of 3 pending connections for the master socket
    if (listen(master_socket, 3) < 0) {
        perror("listen");
        exit(EXIT_FAILURE);
    } 

    // Accept the incoming connection
    addrlen = sizeof(address);
    puts("Waiting for connections ...");

    for (;;) {
        // Clear the socket set
        FD_ZERO(&readfds);

        // Add master socket to set
        FD_SET(master_socket, &readfds);
        max_sd = master_socket;

        // Add child sockets to set
        for (i = 0; i < max_clients; i++) {
            // Socket descriptor
            sd = client_socket[i];

            // If valid socket descriptor then add to read list
            if (sd > 0)
                FD_SET(sd, &readfds);

            // Highest file descriptor number, need it for the select function
            if (sd > max_sd)
                max_sd = sd;
        }

        // Wait for an activity on one of the sockets, timeout is NULL,
        // so wait indefinitely 
        activity = select( max_sd + 1, &readfds, NULL, NULL, NULL);
        if ((activity < 0) && (errno!=EINTR)) {
            printf("select error");
        }

        /*************************/
        /* NEW CLIENT CONNECTION */
        /*************************/

        if (FD_ISSET(master_socket, &readfds)) {
            new_socket = accept(master_socket, (struct sockaddr *)&address, (socklen_t*)&addrlen);
            if (new_socket < 0) {
                perror("accept");
                exit(EXIT_FAILURE);
            }

            // Greetings message
            memset(send_buffer, 0, 1024);
            sprintf(send_buffer, "%s", "greetings");
            size_t res = send(new_socket, send_buffer, strlen(send_buffer), 0);
            if (res != strlen(send_buffer)) {
                perror("send");
            }

            // Add new socket to array of sockets
            for (i = 0; i < max_clients; i++) {
                // If position is empty
                if (client_socket[i] == 0) {
                    client_socket[i] = new_socket;
                    printf(
                        "[ID: %d, IP: %s, PORT: %d]: connected\n",
                        i, inet_ntoa(address.sin_addr), ntohs(address.sin_port)
                    );
                    break;
                }
            }
        }

        /*****************************************/
        /* MESSAGE FROM ALREADY-CONNECTED CLIENT */
        /*****************************************/

        else {
            for (i = 0; i < max_clients; i++) {
                sd = client_socket[i];

                if (FD_ISSET(sd, &readfds)) {
                    // If empty message, closing connection
                    if ((valread = read(sd, buffer, 1024)) == 0) {
                        // Somebody disconnected , get his details and print
                        getpeername(sd, (struct sockaddr*)&address, (socklen_t*)&addrlen);
                        printf(
                            "[ID: %d, IP: %s, PORT: %d]: disconnected\n",
                            i, inet_ntoa(address.sin_addr), ntohs(address.sin_port)
                        );

                        // Close the socket and mark as 0 in list for reuse
                        close(sd);
                        client_socket[i] = 0;
                    } else {
                        // The buffer contains data from client
                        printf(
                            "[ID: %d, IP: %s, PORT: %d]: read \"%s\"\n",
                            i, inet_ntoa(address.sin_addr), ntohs(address.sin_port), buffer
                        );
                        memset(send_buffer, 0, 1024);
                        sprintf(send_buffer, "%s", "message");
                        send(sd, send_buffer, strlen(send_buffer), 0);
                    }
                }
            }
        }
    }

    return 0;
}
