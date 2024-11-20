#include <error.h>
#include <string>

using namespace std;

void log_error(string func_name) {
    cerr << "Error in " << func_name << ": " << SDL_GetError() << endl;
}
