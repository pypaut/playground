#include "client.h"


int main()
{
    /* Variables */
    char *buffer = calloc(1024, sizeof(char));
    char* dir = calloc(1024, sizeof(char));

    Uint64 current_tick = 0;

    /* SDL components */
    SDL_Init(SDL_INIT_VIDEO);

    SDL_Window **window = calloc(1, sizeof(SDL_Window**));
    SDL_Renderer **renderer = calloc(1, sizeof(SDL_Renderer**));

    SDL_Color color;
    color.r = 200;
    color.g = 200;
    color.b = 200;
    color.a = 255;

    player_pos **positions = new_player_pos_list();

    if (SDL_CreateWindowAndRenderer(W * SCALE, H * SCALE, 0, window, renderer)) {
        fprintf(stderr, "%s\n", "error: SDL_CreateWindowAndRenderer\0");
        free(dir);
        free(buffer);
        free(window);
        free(renderer);
        return 1;
    }

    /* Socket components */
    int client_socket_fd;
    init_client(&client_socket_fd);

    /* Main loop */
    for (;;) {
        current_tick = clock_tick(current_tick);

        const Uint8 *keys = SDL_GetKeyboardState(NULL);

        if (check_quit_events(keys)) {
            break;
        }

        float dir_x = 0;
        float dir_y = 0;
        get_dir(keys, &dir_x, &dir_y);

        /* Send to server */
        memset(dir, 0, 1024);
        sprintf(dir, "x:%f,y:%f", dir_x, dir_y);
        send(client_socket_fd, dir, strlen(dir), 0);

        /* Receive from server */
        read(client_socket_fd, buffer, 1023);

        /* Extract positions */
        parse_server_message(positions, buffer);

        /* Draw */
        if (draw(renderer, positions, &color)) {
            break;
        }
    }

    close(client_socket_fd);
    free(dir);
    free(buffer);
    free(window);
    free(renderer);
    free(positions);

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
    if (dt < 1000 / 30) {
        SDL_Delay(1000 / 30 - dt);
    }

    return current_tick;
}

void get_dir(const Uint8 *keys, float *dir_x, float *dir_y) {
    // Go right
    if (keys[SDL_SCANCODE_D]) {
        (*dir_x)++;
    }

    // Go left
    if (keys[SDL_SCANCODE_A]) {
        (*dir_x)--;
    }

    // Go down
    if (keys[SDL_SCANCODE_S]) {
        (*dir_y)++;
    }

    // Go up
    if (keys[SDL_SCANCODE_W]) {
        (*dir_y)--;
    }

    // Normalize direction
    float norm = sqrt(pow(*dir_x, 2) + pow(*dir_y, 2));
    if (norm != 0) {
        *dir_x = *dir_x / norm;
        *dir_y = *dir_y / norm;
    }
}

int check_quit_events(const Uint8 *keys) {
    if (keys[SDL_SCANCODE_ESCAPE]) {
        return 1;
    }

    SDL_Event event;
    while (SDL_PollEvent(&event)) {
        if (event.type == SDL_QUIT) {
            return 1;
        }
    }

    return 0;
}

int draw(SDL_Renderer **renderer, player_pos **positions, SDL_Color *color) {
    /* Background */
    if (SDL_SetRenderDrawColor(*renderer, 0, 0, 0, 255)) {
        fprintf(stderr, "%s\n", "error: SDL_SetRenderDrawColor\0");
        return 1;
    }

    if (SDL_RenderClear(*renderer)) {
        fprintf(stderr, "%s\n", "error: SDL_RenderClear\0");
        return 1;
    }

    /* Players */
    if (SDL_SetRenderDrawColor(
        *renderer,
        color->r,
        color->g,
        color->b,
        color->a)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        return 1;
    }

    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        if (!positions[i]->enabled) {
            continue;
        }

        SDL_Rect rect;
        rect.x = positions[i]->x;
        rect.y = positions[i]->y;
        rect.w = PLAYER_SIZE * SCALE;
        rect.h = PLAYER_SIZE * SCALE;

        SDL_RenderFillRect(*renderer, &rect);
    }


    SDL_RenderPresent(*renderer);
    return 0;
}
