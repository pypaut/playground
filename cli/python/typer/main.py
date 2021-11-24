#!/usr/bin/python3

import typer
import requests


def main():
    app = typer.Typer()

    @app.command()
    def list():
        typer.echo(f"Hello {name}")

    @app.command()
    def goodbye(name: str):
        typer.echo(f"Bye {name}")

    app()


if __name__ == "__main__":
    main()
