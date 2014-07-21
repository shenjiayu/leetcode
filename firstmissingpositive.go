package main

import (
	"fmt"
)

func main() {
	list := []int{5, 6, -1, 3, 1, 2}
	number := fmp(list)
	fmt.Printf("The first missing number is %d", number)
}

func fmp(list []int) int {
	length := len(list)
	smallest := 0
	biggest := 0
	missing := make([]int, 0)
	for i := 0; i < length; i++ {
		if i == 0 {
			smallest = list[i]
			biggest = list[i]
			continue
		}
		if len(missing) > 0 {
			for k, v := range missing {
				if list[i] == v {
					missing = append(missing[:k], missing[k+1:]...)
				}
			}
		} else if len(missing) == 0 && i == length-1 {
			return biggest + 1
		}
		if smallest > list[i] {
			if smallest-list[i] > 1 {
				for j := list[i] + 1; j <= biggest; j++ {
					if j == 0 {
						continue
					}
					missing = append(missing, j)
				}
			}
			smallest = list[i]
			continue
		}
		if list[i] > biggest {
			if list[i]-biggest > 1 {
				for j := biggest + 1; j <= list[i]; j++ {
					if j == 0 {
						continue
					}
					missing = append(missing, j)
				}
			}
			biggest = list[i]
			continue
		}
	}
	return missing[0]
}
