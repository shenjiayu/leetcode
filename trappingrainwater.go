package main

import (
	"fmt"
)

func main() {
	list := []int{2, 1, 0, 2, 1, 0, 1, 3, 2, 1, 3, 1}
	trw(list)
}

func trw(list []int) {
	current := 0
	distance := 0
	potential := 0
	length := len(list)
	for k := 0; k < length; k++ {
		if list[k] == 0 {
			current = k + 1
			continue
		}
		for i := current + 1; i < length; i++ {
			if i == length-1 {
				current++
				potential = 0
				break
			} else if list[i] < list[current] {
				potential += list[current] - list[i]
			} else if list[i] >= list[current] && i-current > 1 {
				distance += potential
				k = i
				current = i
				break
			}
		}
		potential = 0
	}
	fmt.Println(distance)
}
