/*  
 *  hello-1.c - The simplest kernel module.
 */
#include <linux/module.h>	/* Needed by all modules */
#include <linux/kernel.h>	/* Needed for KERN_INFO */

#include "hello-1.h"

MODULE_LICENSE("GPL");
MODULE_AUTHOR("pypaut");
MODULE_DESCRIPTION("my module");

int my_init_module(void) {
	printk(KERN_INFO "Hello world 1.\n");

	/* 
	 * A non 0 return means init_module failed; module can't be loaded.
	 */

	return -1;
}

void my_cleanup_module(void) {
	printk(KERN_INFO "Goodbye world 1.\n");
}

module_init(my_init_module);
module_exit(my_cleanup_module);
