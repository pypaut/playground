obj-m += mykmod.o

all:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules

clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean

dependencies-arch-zen:
	sudo pacman -Sy linux-zen-headers
