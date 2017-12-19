#include <stdio.h>

int main() {
        int a, *p;
	a = 0;
	p = &a;
	printf("%d\n", sizeof(p));
}
