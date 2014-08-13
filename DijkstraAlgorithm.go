package main

import (
	"fmt"
	"math/rand"
)

func randomGraph() [][5]int {
	grid := make([][5]int, 5)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i != j {
				grid[i][j] = rand.Int()%30 + 1
			} else {
				grid[i][j] = 0
			}
		}
	}
	return grid
}

func printGraph(grid [][5]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
}

func da(grid [][5]int) {
	i, j := 0, 0
	for n := 0; n < 5; n++ {
		k := n
		next := 0
		min := 66535
		d := make([]int, 5)
		visited := make([]bool, 5)
		parent := make([]int, 5)
		for i = 0; i < 5; i++ {
			parent[i] = n
			d[i] = grid[k][i]
			if i == k {
				visited[i] = true
			} else {
				visited[i] = false
			}
		}
		for i = 0; i < 5; i++ {
			for j = 0; j < 5; j++ {
				if visited[j] == true {
					continue
				}
				if (d[k] + grid[k][j]) < d[j] {
					d[j] = d[k] + grid[k][j]
					parent[j] = k
				}
				if min > d[j] {
					min = d[j]
					next = j
				}
			}
			k = next
			visited[k] = true
			min = 65535
		}
		fmt.Printf("---------------The starting vertex is %d--------------\n", n)
		for i = 0; i < 5; i++ {
			fmt.Printf("The parent of %d is %d\n", i, parent[i])
			fmt.Printf("The shortest path from v%d to v%d is %d\n", n, i, d[i])
		}
	}
}

func main() {
	//生成一个5X5的矩阵，并且每个方向都带有权值
	grid := randomGraph()
	//打印矩阵
	printGraph(grid)
	//执行算法
	da(grid)
}
