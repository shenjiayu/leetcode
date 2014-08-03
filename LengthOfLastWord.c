#include <stdio.h>

int lolw(char []);

int lolw(char list[]) {
	int count = 0;
	for (int i=0; list[i]!='\0'; i++) {
		if (list[i] == ' ') {
			count = 0;
			continue;
		}
		count++;
	}
	return count;
}

int main() {
	char list[] = "Hello World!";
	int length = lolw(list);
	printf("the length of last word is %d\n", length);
	return 0;
}