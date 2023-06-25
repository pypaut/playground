OBJS = \
	   src/constants.py

all: lint
	python3 main.py

lint:
	black -l 79 main.py ${OBJS}
	flake8 main.py ${OBJS}

