package main

import "fmt"

func main() {
	fmt.Println(half(7))

	fmt.Println(fib(5))
}

func sum(slice []int) int {
	return 0
}

func half(a int) bool {
	return (a/2)%2 == 0
}

func greatest(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) int {
	fmt.Println(n)
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
