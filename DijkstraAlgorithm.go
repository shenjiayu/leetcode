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
				grid[i][j] = rand.Int()%50 + 1
				grid[j][i] = grid[i][j]
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
	k, i, j, next, min := 0, 0, 0, 0, 65535
	d := make([]int, 5)
	visited := make([]bool, 5)
	for i = 0; i < 5; i++ {
		d[i] = grid[0][i]
		if i == k {
			visited[i] = true
		} else {
			visited[i] = false
		}
	}
	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			if visited[j] == true || j == k {
				continue
			}
			if (d[k] + grid[k][j]) < d[j] {
				d[j] = d[k] + grid[k][j]
			}
			if min > d[j] {
				min = d[j]
				next = j
			}
		}
		min = 65535
		if next == k {
			k = 0
		}
		fmt.Printf("The Parent of %d is %d\n", next, k)
		k = next
		visited[k] = true
		fmt.Printf("The shortest path from v0 to v%d is %d\n", i, d[i])
	}
}

func main() {
	//生成一个5X5的矩阵，并且每个点都带有权值
	grid := randomGraph()
	//打印矩阵
	printGraph(grid)
	//执行算法
	da(grid)
}
