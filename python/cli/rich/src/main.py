import click

from rich import print


@click.command()
def main():
    print("Hello, world!")


if __name__ == "__main__":
    main()
