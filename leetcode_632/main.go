package main

import "fmt"

func smallestRange(nums [][]int) []int {
	base := 100000
	k := len(nums)
	pre := make([][]int, k)
	next := make([][]int, k)
	element := make([][]int, 2*base+1)
	for ith, vs := range nums {
		lx := -1
		for i := range vs {
			if i > 0 && vs[i] == vs[i-1] {
				pre[ith] = append(pre[ith], pre[ith][i-1])
				continue
			}
			tv := vs[i] + base
			element[tv] = append(element[tv], ith)
			pre[ith] = append(pre[ith], lx)
			lx = tv
		}
		rx := 2 * base + 1
		next[ith] = make([]int, len(vs))
		for i := len(vs) - 1; i >= 0; i-- {
			if i < len(vs)-1 && vs[i] == vs[i+1] {
				next[ith][i] = next[ith][i+1]
				continue
			}
			tv := vs[i] + base
			next[ith][i] = rx
			rx = tv
		}
	}
	ansL := 0
	ansR := 2 * base
	r := 0
	cnt := 0
	pos := make([]int, k)
	spos := make([]int, k)
	for i := 0; i <= 2*base; i++ {
		for ; r <= 2*base; r++ {
			if cnt == k {
				break
			}
			for _, ith := range element[r] {
				index := pos[ith]
				if pre[ith][index] < i {
					cnt++
			//		fmt.Println(ith, r-base, index, cnt)
				}
				for index < len(nums[ith]) && nums[ith][index] == r-base {
					index++
				}
				pos[ith] = index
			}
		}
		if cnt == k {
			if  r-i < ansR-ansL+1 {
				ansL = i
				ansR = r - 1
			}
		}
		for _, ith := range element[i] {
			index := spos[ith]
			//fmt.Println(next[ith][index]-base, ith, index)
			if next[ith][index] >= r {
			//	fmt.Println(i-base, r-base, index)
				cnt--
			}
			for index < len(nums[ith]) && nums[ith][index] == i-base {
				index++
			}
			spos[ith] = index
		}
	}
	return []int{ansL - base, ansR - base}
}
func main() {
	var nums [][]int
	a := []int{4, 10, 15, 24, 26}
	b := []int{0, 9, 12, 20}
	c := []int{5, 18, 22, 30}
	nums = append(nums, a)
	nums = append(nums, b)
	nums = append(nums, c)
	fmt.Println(smallestRange(nums))
	a = []int{10, 10}
	b = []int{11, 11}
	nums = nums[:0]
	nums = append(nums, a)
	nums = append(nums, b)
	fmt.Println(smallestRange(nums))

}
