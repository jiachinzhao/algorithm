package main

import (
	"fmt"
)
// 首先考虑超过n/2次数，注意到最多只有一个，若存在，其他数字的总和会比这个数字的数量少，可以利用这一点求出这个数字
// 同理超过n/3次数的数字最多只有两个，也可以利用类似的性质
func majorityElement(nums []int) []int {
	var ans []int
	mx := []int{-1, -1}
	cntMx := []int{0, 0}

	for _, v := range nums {
		if mx[0] == v {
			cntMx[0]++
		} else if mx[1] == v {
			cntMx[1]++
		} else if cntMx[0] == 0 {
			mx[0] = v
			cntMx[0] = 1
		} else if cntMx[1] == 0 {
			mx[1] = v
			cntMx[1] = 1
		} else {
			cntMx[0]--
			cntMx[1]--
		}
	}
	check := func(x int) bool {
		count := 0
		for _, v := range nums {
			if v == x {
				count++
			}
		}
		return count > len(nums)/3
	}
	for _, v := range mx {
		if check(v) {
			ans = append(ans, v)
		}
	}

	return ans

}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
	fmt.Println(majorityElement([]int{1, 2, 3}))
}
