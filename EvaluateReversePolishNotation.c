#include <stdio.h>

void erpn(char [], int len);

void erpn(char list[], int len){
	int ans = 0;
	int current = 1;
	for (int i=0; i < len; i++){
		switch (list[i]) {
			case '+':
				ans = list[0] + list[current];
				list[0] = ans;
				current = 0;
				break;
			case '-':
				ans = list[0] - list[current];
				list[0] = ans;
				current = 0;
				break;
			case '*':
				ans = list[0] * list[current];
				list[0] = ans;
				current = 0;
				break;
			case '/':
				ans = list[0] / list[current];
				list[0] = ans;
				current = 0;
				break;
			default:
				current = i;
		}
		list[i] -= 48;
	}
	printf("%d\n", ans);
}

int main(){
	char list[] = {'4', '13', '+', '3', '*'};
	int len = sizeof(list) / sizeof(list[0]);
	erpn(list, len);
	return 0;
}