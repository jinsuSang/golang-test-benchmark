package main

import "fmt"

func Square(x int) int {
	return x * x
}

func Fibo(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return b
}

func FiboRecursive(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	return FiboRecursive(n-1) + FiboRecursive(n-2)
}

func Dfs1(v int, visited []bool, graph [][]int) {
	stack := []int{v}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[node] = true
		// fmt.Println(node)
		for _, nd := range graph[node] {
			if visited[nd] == false {
				stack = append(stack, nd)
			}
		}
	}
}

func Dfs2(v int, visited []bool, graph [][]int) {
	visited[v] = true
	// fmt.Println(v)
	for _, node := range graph[v] {
		if visited[node] == false {
			Dfs2(node, visited, graph)
		}
	}
}

func main() {
	fmt.Printf("9 * 9 = %d\n", Square(9))

	fmt.Println(Fibo(5))
	fmt.Println(FiboRecursive(5))
}
