#include "client.h"

int main()
{
    char buffer[1024] = { 0 };
    char* dir = calloc(1024, sizeof(char));

    Uint64 current_tick = 0;

    int client_socket_fd;
    init_client(&client_socket_fd);

    /* SDL components */
    SDL_Init(SDL_INIT_VIDEO);

    SDL_Window **window = calloc(1, sizeof(SDL_Window**));
    SDL_Renderer **renderer = calloc(1, sizeof(SDL_Renderer**));

    int W = 1920 / 2;
    int H = 1080 / 2;

    if (SDL_CreateWindowAndRenderer(W, H, 0, window, renderer)) {
        fprintf(stderr, "%s\n", "error: SDL_CreateWindowAndRenderer\0");
        close(client_socket_fd);
        free(dir);
        free(window);
        free(renderer);
        return 1;
    }

    /* Main loop */
    for (;;) {
        current_tick = clock_tick(current_tick);

        // Check quit event
        SDL_Event event;
        while (SDL_PollEvent(&event)) {
            if (event.type == SDL_QUIT) {
                break;
            }
        }

        // Read keys events
        const Uint8 *keys = SDL_GetKeyboardState(NULL);

        // Escape
        if (keys[SDL_SCANCODE_ESCAPE]) {
            break;
        }

        float dir_x = 0;
        float dir_y = 0;

        // Go right
        if (keys[SDL_SCANCODE_D]) {
            dir_x++;
        }

        // Go left
        if (keys[SDL_SCANCODE_A]) {
            dir_x--;
        }

        // Go down
        if (keys[SDL_SCANCODE_S]) {
            dir_y++;
        }

        // Go up
        if (keys[SDL_SCANCODE_W]) {
            dir_y--;
        }

        // Send direction/input to server
        // memset(dir, 0, len(dir));
        sprintf(dir, "x:%f,y:%f", dir_x, dir_y);
        send(client_socket_fd, dir, strlen(dir), 0);
        read(client_socket_fd, buffer, 1024 - 1);
        printf("%s\n", buffer);

        // Draw
        if (SDL_SetRenderDrawColor(*renderer, 0, 0, 0, 255)) {
            fprintf(stderr, "%s\n", "error: SDL_SetRenderDrawColor\0");
            close(client_socket_fd);
            free(dir);
            free(window);
            free(renderer);
            break;
        }

        if (SDL_RenderClear(*renderer)) {
            fprintf(stderr, "%s\n", "error: SDL_RenderClear\0");
            close(client_socket_fd);
            free(dir);
            free(window);
            free(renderer);
            break;
        }

        // Final render
        SDL_RenderPresent(*renderer);
    }

    close(client_socket_fd);
    free(dir);
    free(window);
    free(renderer);

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
