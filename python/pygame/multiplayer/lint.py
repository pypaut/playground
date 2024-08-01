import os


def main():
    print("ruff format...")
    os.system("poetry run ruff format")

    print("\nruff check...")
    os.system("poetry run ruff check")


if __name__ == "__main__":
    main()
