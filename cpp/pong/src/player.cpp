#include <SDL2/SDL.h>
#include <error.h>
#include <player.h>


Player::Player() {}

Player::~Player() {
    delete(this->rect);
}

void Player::SetKeys(Uint8 up_key, Uint8 down_key) {
    this->up_key = up_key;
    this->down_key = down_key;
}

void Player::SetColor(Uint8 r, Uint8 g, Uint8 b, Uint8 a) {
    this->color_r = r;
    this->color_g = g;
    this->color_b = b;
    this->color_a = a;
}

void Player::SetRect(int x, int y, int w, int h) {
    this->rect = new SDL_Rect{x, y, w, h};
}

void Player::Update(const Uint8 *keys, int win_height) {
    if (keys[this->up_key]) {
      this->MoveUp();
    }

    if (keys[this->down_key]) {
      this->MoveDown(win_height);
    }
}

void Player::MoveUp() {
    if (this->rect->y > 0) {
        this->rect->y--;
    }
}

void Player::MoveDown(int window_height) {
    if (this->rect->y + this->rect->h < window_height) {
        this->rect->y++;
    }
}

int Player::Draw(SDL_Renderer *renderer) {
    if (SDL_SetRenderDrawColor(
                renderer,
                this->color_r,
                this->color_g,
                this->color_b,
                this->color_a
    )) {
        log_error("Player::Draw::SDL_SetRenderDrawColor");
        return 1;
    }

    if (SDL_RenderFillRect(renderer, this->rect)) {
        log_error("Player::Draw::SDL_RenderFillRect");
        return 1;
    }

    return 0;
}
