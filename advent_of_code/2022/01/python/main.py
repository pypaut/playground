#!/usr/bin/python


def main():
    elves = [0]
    lines = []
    current_i = 0

    with open("../input", "r") as f:
        lines = f.readlines()
    
    for l in lines:
        if l in {"", " ", "\n"}:
            current_i += 1
            elves.append(0)
        else:
            elves[current_i] += int(l)

    
    print(max(elves))


if __name__ == "__main__":
    main()
