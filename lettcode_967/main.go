package main

import (
	"fmt"
	"sync"
)

func dfs(x, num, cur, N, K int, ch chan int) {
	if cur == 0 && x == 0 && N > 1 {
		return
	}
	if cur == N-1 {
		ch <- num
		return
	}
	if x+K < 10 {
		dfs(x+K, num * 10 + x + K,cur+1, N, K, ch)
	}
	if K != 0 && x-K >= 0 {
		dfs(x+K, num * 10 + x - K,cur+1, N, K, ch)
	}
}
func numsSameConsecDiff(N int, K int) []int {
	var ans []int
	ansCh := make(chan int)
	var sc sync.WaitGroup
	for i := 0; i <= 9; i++ {
		sc.Add(1)
		go func(i int) {
			dfs(i, i,0, N, K, ansCh)
			sc.Done()
		}(i)
	}

	go func() {
		sc.Wait()
		close(ansCh)
	}()

	for x := range ansCh {
		ans = append(ans, x)
	}

	return ans

}
func main() {
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(9, 2))
}
