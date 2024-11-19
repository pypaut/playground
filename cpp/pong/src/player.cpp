#include <SDL2/SDL.h>
#include <error.h>
#include <player.h>

#include <algorithm>


Player::Player() {}

Player::~Player() {
    delete(this->rect);
}

void Player::SetKeys(Uint8 up_key, Uint8 down_key) {
    this->up_key = up_key;
    this->down_key = down_key;
}

void Player::SetSpeed(int speed) {
    this->speed = speed;
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

void Player::Update(const Uint8 *keys, int win_height, Uint64 dt) {
    if (keys[this->up_key]) {
        this->rect->y -= this->speed * dt;
    }

    if (keys[this->down_key]) {
        this->rect->y += this->speed * dt;
    }

    this->rect->y = std::clamp(this->rect->y, 0, win_height - this->rect->h);
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
