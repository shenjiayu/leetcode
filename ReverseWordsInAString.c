#include <stdio.h>

void rwias(char [],int len);

void rwias(char list1[],int len){
	char list2[16] = {" "};
	int step = 0;
	for (int i = 0; i < len; i++){
		if (list1[i] == '\0'){
			for (int j=len-i; step > 0; j++){
				list2[j] = list1[i-step];
				step--;
			}
		}else if (list1[i] == ' '){
			list2[len-i-1] = ' ';
			for (int j=len-i; step > 0; j++){
				list2[j] = list1[i-step];
				step--;
			}
			i++;
		}
		step++;
	}
	printf("%s\n", list2);
}

int main(){
	char list1[] = "the sky is blue and clean";
	int len = sizeof(list1) / sizeof(list1[0]);
	rwias(list1, len);
	return 0;
}