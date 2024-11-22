#include <SDL2/SDL.h>
#include <ball.h>
#include <cmath>
#include <error.h>
#include <random>

Ball::Ball() {}

Ball::~Ball() { delete (this->rect); }

void Ball::SetSpeed(int speed) { this->speed = speed; }

void Ball::SetColor(Uint8 r, Uint8 g, Uint8 b, Uint8 a) {
    this->color_r = r;
    this->color_g = g;
    this->color_b = b;
    this->color_a = a;
}

void Ball::SetRect(int x, int y, int w, int h) {
    this->rect = new SDL_Rect{x, y, w, h};
}

void Ball::SetDir(float x, float y) {
    this->dir_x = x;
    this->dir_y = y;
}

float Ball::GetDirX() { return this->dir_x; }

float Ball::GetDirY() { return this->dir_y; }

bool Ball::Update(int win_width, int win_height, Uint64 dt) {
    this->NormalizeDir();

    // Update position
    this->rect->x = this->rect->x + this->dir_x * this->speed * dt;
    this->rect->y = this->rect->y + this->dir_y * this->speed * dt;

    // Check for vertical walls collision
    if (this->rect->x < 0 || this->rect->x + this->rect->w > win_width) {
        return false;
    }

    // Check for horizontal walls collision
    // Two different conditions to avoid changing dir multiple times while
    // the ball is still verifying this condition
    if (this->rect->y < 0) {
        this->dir_y = abs(this->dir_y);
    }
    if (this->rect->y + this->rect->h > win_height) {
        this->dir_y = -abs(this->dir_y);
    }

    return true;
}

int Ball::Draw(SDL_Renderer *renderer) {
    if (SDL_SetRenderDrawColor(renderer, this->color_r, this->color_g,
                               this->color_b, this->color_a)) {
        log_error("Ball::Draw::SDL_SetRenderDrawColor");
        return 1;
    }

    if (SDL_RenderFillRect(renderer, this->rect)) {
        log_error("Ball::Draw::SDL_RenderFillRect");
        return 1;
    }

    return 0;
}

void Ball::Build(int win_width, int win_height) {
    int ball_side = win_width / 100;
    int ball_x = (win_width - ball_side) / 2;
    int ball_y = (win_height - ball_side) / 2;

    this->SetRect(ball_x, ball_y, ball_side, ball_side);
    this->SetColor(255, 255, 255, 255);
    this->speed = 0.5;

    // Random initial ball direction
    std::random_device r;
    std::default_random_engine e1(r());

    std::uniform_int_distribution<int> uniform_dist_x(-80, 80);
    std::uniform_int_distribution<int> uniform_dist_y(-4, 4);

    int x = uniform_dist_x(e1);
    int y = uniform_dist_y(e1);

    this->dir_x = x;
    this->dir_y = y;
}

void Ball::NormalizeDir() {
    float norm = sqrt(pow(this->dir_x, 2) + pow(this->dir_y, 2));

    this->dir_x = this->dir_x / norm;
    this->dir_y = this->dir_y / norm;
}
