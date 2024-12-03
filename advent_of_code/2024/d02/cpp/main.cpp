#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>

using namespace std;

vector<int[]> *ParseLevels(string input_file) {
    vector<int[]> *levels = new vector<int[]>;

    ifstream input(input_file);

    string line;
    while (getline(input, line)) {
        istringstream line_reader(line);

        do {
            string token;
            line_reader >> token;
            int new_value = stoi(token);

            new_level->push_back(new_value);
        } while (line_reader);

        levels->push_back(*new_level);
    }

    input.close();
    return levels;
}

int main() {
    vector<vector<int>> *levels = ParseLevels("input_test");
    for (vector<int> level : levels) {
        for (int v : level) {
            cout << v << "test";
        }

        // cout << endl;
    }

    return 0;
}
