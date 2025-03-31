/*  
 *  mykmod.c - The simplest kernel module.
 */
#include <linux/module.h>	/* Needed by all modules */
#include <linux/kernel.h>	/* Needed for KERN_INFO */
#include <linux/init.h>		/* Needed for the macros */

MODULE_LICENSE("GPL");
MODULE_AUTHOR("pypaut");
MODULE_DESCRIPTION("my kernel module");

static int mykmod_data __initdata = 2;

static int __init mykmod_init(void) {
	printk(KERN_INFO "hello, kernel! %d\n", mykmod_data);
	return 0;
}

static void __exit mykmod_exit(void) {
	printk(KERN_INFO "bye, kernel!\n");
}

module_init(mykmod_init);
module_exit(mykmod_exit);
