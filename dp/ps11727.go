package main

import (
	"fmt"
	"os"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)

	if T == 1 {
		fmt.Printf("%d", 1)
		os.Exit(0)
	}

	dp := make([]int, 1001)
	dp[1] = 1
	dp[2] = 3

	for i := 3; i <= T; i++ {
		dp[i] = (dp[i-2] * 2 + dp[i-1]) % 10007
	}

	fmt.Printf("%d", dp[T])
}
