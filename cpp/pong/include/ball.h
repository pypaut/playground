#include <SDL2/SDL.h>

class Ball {
  public:
    Ball();
    ~Ball();

    void Build(int win_width, int win_height);

    void SetColor(Uint8 r, Uint8 g, Uint8 b, Uint8 a);
    void SetSpeed(int speed);
    void SetRect(int x, int y, int w, int h);
    void SetDir(float x, float y);

    float GetDirX();
    float GetDirY();

    bool Update(int win_width, int win_height, Uint64 dt);
    int Draw(SDL_Renderer *renderer);

    void NormalizeDir();

    SDL_Rect *rect;

  private:
    float speed;

    float dir_x;
    float dir_y;

    Uint8 color_r;
    Uint8 color_g;
    Uint8 color_b;
    Uint8 color_a;
};
