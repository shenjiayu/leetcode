package main

import (
	"fmt"
)

const (
	None = -1
)

func main() {
	list := []int{3, 9, 20, None, None, 15, 7, 10}
	order := btlot2(list)
	fmt.Println(order)
}

func btlot2(list []int) [][2]int {
	length := len(list)
	order := make([][2]int, length/2)
	node := 0
	for i := length - 1; i >= 0; i -= 2 {
		if i == 0 {
			order[node][0] = list[i]
			break
		} else if list[i] != None && list[i-1] != None {
			order[node][0] = list[i-1]
			order[node][1] = list[i]
		} else if list[i] != None && list[i-1] == None {
			order[node][0] = list[i]
			order[node][1] = -1
		} else if list[i-1] != None && list[i] == None {
			order[node][0] = list[i-1]
			order[node][1] = -1
		} else {
			continue
		}
		node++
	}
	return order
}
