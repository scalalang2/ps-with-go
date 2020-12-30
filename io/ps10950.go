package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {
		var x, y int64
		fmt.Scanf("%d %d", &x, &y)

		fmt.Println(x + y)
	}
}