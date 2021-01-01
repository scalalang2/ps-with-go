package main

import "fmt"

func main() {
	var mod int64
	mod = 1000000000

	var T int64
	fmt.Scanf("%d", &T)

	var dp [101][10]int64
	for i := 1; i < 10; i++ {
		dp[1][i] = 1
	}

	for i := 2; int64(i) <= T; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j+1]
			}  else if j == 9 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = (dp[i-1][j-1] + dp[i-1][j+1]) % mod
			}
		}
	}

	var sum int64
	for i := 0; i < 10; i++ {
		sum += dp[T][i]
	}
	sum = sum % mod
	fmt.Println(sum)
}
