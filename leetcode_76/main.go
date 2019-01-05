package main

import "fmt"

func minWindow(s string, t string) string {
	tCnt := make([]int, 256)
	difCnt := 0
	for _, v := range t {
		if tCnt[v] == 0 {
			difCnt++
		}
		tCnt[v]++
	}
	mapCnt := make([]int, 256)
	cnt := 0
	r := 0
	ansInx := -1
	ansLen := len(s) + 1
	for i, v := range s {
		for ; r < len(s) && cnt != difCnt; r++ {
			if tCnt[s[r]] == 0 {
				continue
			}
			mapCnt[s[r]]++
			if mapCnt[s[r]] == tCnt[s[r]] {
				cnt++
			}
		}
		if cnt == difCnt {
			if r-i < ansLen {
				ansLen = r - i
				ansInx = i
			}
		}
		if tCnt[v] > 0 {
			if mapCnt[v] == tCnt[v] {
				cnt--
			}
			mapCnt[v]--
		}
	}
	if ansInx == -1 {
		return ""
	}
	//fmt.Println(ansInx, ansLen)
	return s[ansInx : ansInx+ansLen]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
	s = "aabbc"
	t = "aac"
	fmt.Println(minWindow(s, t))
}
