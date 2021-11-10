#!/usr/bin/python3

import uvicorn

from fastapi import FastAPI
from foobarqix import compute

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/foobarqix/{input_str}")
def read_item(input_str: str):
    return {"message": compute(input_str)}


if __name__ == "__main__":
    uvicorn.run(app, host="localhost", port=8000)
