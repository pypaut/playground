#include <algorithm>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

int main() {
    vector<int> elves = {0};
    size_t current_i = 0;

    string line;
    ifstream input("../input");
    while (getline(input, line)) {
        if (!line.compare(" ") || !line.compare("\n") || !line.compare("")) {
            current_i++;
            elves.push_back(0);
        } else {
            elves[current_i] = elves.at(current_i) + stoi(line, nullptr, 10);
        }
    }
    input.close();

    double sum_cal = 0;

    for (size_t _ = 0; _ < 3; _++) {
        std::vector<int>::iterator max_iter =
            max_element(elves.begin(), elves.end());
        sum_cal = sum_cal + *max_iter;

        double max_i = distance(elves.begin(), max_iter);
        elves[max_i] = 0;
    }

    cout << sum_cal << endl;
    return 0;
}
