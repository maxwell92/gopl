#include <stdio.h>

void main() {
	int a = 1;	
	switch(a) {
		case 1: {
			printf("this is case1\n");
  			break;
		}
		case 2: {
			printf("this is case2\n");
		}
		default: {
			printf("this is default\n");
		}
	}
}
