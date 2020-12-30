package main

import "fmt"

func main() {
	for {
		var x, y int
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil { break }

		fmt.Println(x + y)
	}
}
