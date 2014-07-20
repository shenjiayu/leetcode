package main

import (
	"fmt"
)

func main() {
	list := []int{2, 1, 0, 2, 1, 0, 1, 3, 2, 1, 3, 1}
	test(list)
}

func test(list []int) {
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
			if list[i] < list[current] {
				potential += list[current] - list[i]
				if i == length-1 {
					current++
					potential = 0
					break
				}
			} else if list[i] >= list[current] && i-current > 1 {
				fmt.Printf("current node is %d\n", current)
				fmt.Printf("node %d\n", i)
				fmt.Printf("potential %s\n", potential)
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
