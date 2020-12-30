package main

import "fmt"

func main() {
	for {
		var x, y int
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil { break }
		if x == 0 && y == 0 { break }
		fmt.Println(x + y)
	}
}
