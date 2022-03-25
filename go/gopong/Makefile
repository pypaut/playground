SHELL := /bin/bash

all:
	go build .

install_dependencies:
	sudo apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev
	go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}
