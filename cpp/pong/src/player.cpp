#include <SDL2/SDL.h>
#include <error.h>
#include <player.h>

#include <algorithm>

Player::Player() {}

Player::~Player() { delete (this->rect); }

void Player::BuildDefaultPlayer1(int win_height, int win_width) {
    this->SetColor(255, 255, 255, 255);

    int player_w = win_width / 100;
    int player_x = win_width / 20;

    int player_h = win_height / 5;
    int player_y = (win_height - player_h) / 2;

    this->SetRect(player_x, player_y, player_w, player_h);
    this->SetKeys(SDL_SCANCODE_W, SDL_SCANCODE_S);
    this->SetSpeed(1);
}

void Player::BuildDefaultPlayer2(int win_height, int win_width) {
    this->SetColor(255, 255, 255, 255);

    int player_w = win_width / 100;
    int player_x = win_width - (win_width / 20) - player_w;

    int player_h = win_height / 5;
    int player_y = (win_height - player_h) / 2;

    this->SetRect(player_x, player_y, player_w, player_h);
    this->SetKeys(SDL_SCANCODE_UP, SDL_SCANCODE_DOWN);
    this->SetSpeed(1);
}

void Player::SetKeys(Uint8 up_key, Uint8 down_key) {
    this->up_key = up_key;
    this->down_key = down_key;
}

void Player::SetSpeed(int speed) { this->speed = speed; }

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
    if (SDL_SetRenderDrawColor(renderer, this->color_r, this->color_g,
                               this->color_b, this->color_a)) {
        log_error("Player::Draw::SDL_SetRenderDrawColor");
        return 1;
    }

    if (SDL_RenderFillRect(renderer, this->rect)) {
        log_error("Player::Draw::SDL_RenderFillRect");
        return 1;
    }

    return 0;
}

float Player::GetCenterY() {
    return this->rect->y + (this->rect->h/2);
}
