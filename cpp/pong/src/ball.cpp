#include <SDL2/SDL.h>
#include <ball.h>
#include <cmath>
#include <error.h>
#include <random>


Ball::Ball() {}

Ball::~Ball() {
    delete(this->rect);
}

void Ball::SetSpeed(int speed) {
    this->speed = speed;
}

void Ball::SetColor(Uint8 r, Uint8 g, Uint8 b, Uint8 a) {
    this->color_r = r;
    this->color_g = g;
    this->color_b = b;
    this->color_a = a;
}

void Ball::SetRect(int x, int y, int w, int h) {
    this->rect = new SDL_Rect{x, y, w, h};
}

void Ball::Update(int win_width, int win_height, Uint64 dt) {
    this->NormalizeDir();

    this->rect->x = this->rect->x + this->dir_x*this->speed*dt;
    this->rect->y = this->rect->y + this->dir_y*this->speed*dt;

    win_width = win_width;
    win_height = win_height;
}

int Ball::Draw(SDL_Renderer *renderer) {
    if (SDL_SetRenderDrawColor(
                renderer,
                this->color_r,
                this->color_g,
                this->color_b,
                this->color_a
    )) {
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
    std::uniform_int_distribution<int> uniform_dist_y(-80, 80);

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
