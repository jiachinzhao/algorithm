package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("", "a"))
	fmt.Println(minDistance("a", "ab"))
}

//we define dp[i][j] is -min distance -> w1[...i] and w2[...j] match each other,
//we consider w1[i] and w2[j]
//if w1[i] = w2[j], it can change to dp[i - 1][j - 1]
//otherwise, we can use replace or insert(equal to delete) change to {dp[i-1][j-1], dp[i-1][j], dp[i][j - 1]}

func minDistance(word1 string, word2 string) int {
	if len(word2) > len(word1) {
		word1, word2 = word2, word1
	}
	if len(word2) == 0 {
		return len(word1)
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	dp := make([][]int, len(word1))
	for i := range word1 {
		dp[i] = make([]int, len(word2))
	}

	for i, w1 := range word1 {
		for j, w2 := range word2 {
			if w1 == w2 {
				if i == 0 {
					dp[i][j] = j
				} else if j == 0 {
					dp[i][j] = i
				} else {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				if i == 0 && j == 0 {
					dp[i][j] = 1
				} else if j == 0 {
					dp[i][j] = dp[i-1][j] + 1
				} else if i == 0 {
					dp[i][j] = dp[i][j-1] + 1
				} else {
					dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				}
			}
		}
	}
	return dp[len(word1)-1][len(word2)-1]
}
