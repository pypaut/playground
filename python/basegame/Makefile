OBJS = \
	   src/block.py \
	   src/constants.py \
	   src/player.py \
	   src/sprites.py

all: lint
	python3 main.py

debug: lint
	python3 main.py debug

lint:
	black -l 79 main.py ${OBJS}
	flake8 main.py ${OBJS}

