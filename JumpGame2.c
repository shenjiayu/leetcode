#include <stdio.h>

int jg2(int [], int);

int max(int, int);

int jg2(int list[], int length) {
	int last = 0;
	int ret = 0;
	int curr = 0;
	if (list[0] > length) {
		ret = 1;
		return ret;
	}
	for (int i=0; i<length; i++) {
		if (i > last) {
			if (last == curr) {
				return -1;
			}
			last = curr;
			ret++;
		}
		curr = max(curr, i+list[i]);
	}
	return ret;
}

int max(int curr, int number) {
	if (curr>number) {
		return curr;
	}
	return number;
}

int main() {
	int list[] = {2, 1, 2, 1, 4};
	int length = sizeof(list) / sizeof(list[0]);
	int min = jg2(list, length);
	printf("%d\n", min);
	return 0;
}