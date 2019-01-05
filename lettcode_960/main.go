package main

import "fmt"

func minDeletionSize(A []string) int {
	check  := func(x, y int) bool {
		for _, s := range A{
			if s[x] > s[y]{
				return false
			}
		}
		return true
	}
	eleLen := len(A[0])
	dp := make([]int, eleLen)
	ans := 1
	for i := 0;i < eleLen;i++{
		dp[i] = 1
		for j := 0;j < i;j++{
			if check(j,i){
				if dp[j] + 1 > dp[i]{
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > ans{
			ans = dp[i]
		}
	}
	return eleLen - ans
}
func main() {

	A := []string{
		"babca","bbazb",
	}
	fmt.Println(minDeletionSize(A))
}
