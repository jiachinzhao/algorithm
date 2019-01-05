package main

import (
	"fmt"
)

func minDeletionSize(A []string) int {
	if len(A) == 0 {
		return 0
	}
	ans := 0
	ls := len(A[0])
	vis := make(map[int]bool)
	for i := 0; i < ls; i++ {
		flag := false
		for j := 1; j < len(A); j++ {
			if !vis[j] && A[j][i] < A[j-1][i] {
				flag = true
				break
			}
		}

		if flag {
			ans++
		} else {
			for j := 1; j < len(A); j++ {
				if !vis[j] && A[j][i] > A[j-1][i] {
					vis[j] = true
				}
			}
		}
	}
	return ans
}
func main() {

	var A []string
	A = append(A, "zyx")
	A = append(A, "wvu")
	A = append(A, "tsr")
	fmt.Println(minDeletionSize(A))
	A = A[:0]
	A = append(A, "ca")
	A = append(A, "bb")
	A = append(A, "ac")
	fmt.Println(minDeletionSize(A))
	A = A[:0]
	A = append(A, "bbjwefkpb")
	A = append(A, "axmksfchw")
	fmt.Println(minDeletionSize(A))
}
