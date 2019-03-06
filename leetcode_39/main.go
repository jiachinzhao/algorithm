package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	var ans [][]int
	var plan []int
	ch := make(chan []int)
	go func() {
		dfs(0, 0, target, candidates, plan, ch)
		close(ch)
	}()
	for {
		if v, ok := <-ch; ok {
			ans = append(ans, v)
		} else {
			break
		}
	}
	return ans
}
func dfs(cur, sum, target int, candidates, plan []int, ch chan []int) {
	if sum == target {
		dst := make([]int, len(plan))
		copy(dst, plan)
		ch <- dst
		return
	}
	if cur == len(candidates) {
		return
	}
	if cur > 0 && candidates[cur] == candidates[cur-1] {
		dfs(cur+1, sum, target, candidates, plan, ch)
		return
	}
	v := candidates[cur]
	cnt := 0
	for {
		dfs(cur+1, sum, target, candidates, plan, ch)
		sum += v
		if sum > target {
			break
		}
		cnt++
		plan = append(plan, v)
	}
	plan = plan[:len(plan)-cnt]
}

func main() {
	arr := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(arr, 7))

}
