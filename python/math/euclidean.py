#!/usr/bin/python3


def euclidean(a, b):
    q = a // b
    r = a % b

    while True:
        a = b
        b = r
        q = a // b
        next_r = a % b
        if next_r == 0:
            break
        r = next_r

    return r


def rec_euclidean(a, b):
    q = a // b
    r = a % b
    if r == 0:
        return b

    return rec_euclidean(b, r)


def rec_extended_euclidean(a, b):
    if a == 0:
        return b, 0, 1

    gcd, u1, u1 = rec_extended_euclidean(b % a, a)

    u = v1 - (b // a) * u1
    v = u1

    return gcd, u, v


def main():
    print(euclidean(243, 198))
    print(rec_euclidean(243, 198))
    print(rec_extended_euclidean(243, 198))


if __name__ == "__main__":
    main()
