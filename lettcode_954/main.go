package main

import (
	"fmt"
	"sort"
)

func canReorderDoubled(A []int) bool {
	check := func(a []int) bool {
		sort.Ints(a)
		vis := make(map[int]bool)
		cur := 1
		cnt := 0
		for index, v := range a {
			if cur <= index{
				cur = index + 1
			}
			for ; cur < len(a); cur++ {
				if vis[index] {
					break
				}
				if vis[cur] || a[cur] < 2 * v{
					continue
				}
				if a[cur] > 2*v {
					break
				}
				vis[index] = true
				vis[cur] = true
				cnt += 2
			}
		}
		fmt.Println(cnt)
		return cnt == len(a)
	}
	neg := make([]int, 0, len(A))
	pos := make([]int, 0, len(A))
	for _, v := range A {
		if v < 0 {
			neg = append(neg, -v)
		} else {
			pos = append(pos, v)
		}
	}
	return check(neg) && check(pos)
}
func main() {
	a := []int{0, 0, 0, 0}

	fmt.Println(canReorderDoubled(a))
}
