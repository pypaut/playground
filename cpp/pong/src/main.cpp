#include <error.h>
#include <game.h>

using namespace std;

int main() {
    int W = 1920;
    int H = 1080;

    Game g = Game(W, H);
    return g.Run();
}
