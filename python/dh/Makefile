OBJS = \
	   src/constants.py \
	   src/game.py \
	   src/menu.py

all: lint
	python3 main.py

lint:
	black -l 79 main.py ${OBJS}
	flake8 main.py ${OBJS}

