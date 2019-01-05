package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	var ans []int
	ls := len(s)
	lw := len(words)
	wlen := 0

	for _, v := range words {
		wlen += len(v)
	}
	find := func(x uint8, match []int) bool {
		parMatch := false
		oneMatch := false
		for i, w := range words {
			ll := len(w)
			if  ll == match[i] {
				continue
			}
			if match[i] > 0 {
				parMatch = true
				if x == w[match[i]] {
					oneMatch = true
					match[i]++
					if match[i] == ll{
						for j, ww := range words{
							if j != i && match[j] < len(ww)	&& match[j] > 0{
								match[j] = 0
							}
						}
						return true
					}
				} else{
					match[i] = 0
				}
			}
		}
		if parMatch{
			return oneMatch
		}
		parMatch = false
		for i, w := range words{
			if match[i] == 0 && w[match[i]] == x{
				match[i]++
				if match[i] == len(w){
						for j, ww := range words{
							if j != i && match[j] < len(ww)	&& match[j] > 0{
								match[j] = 0
							}
						}
						return true
					}
				parMatch = true
			}
		}
		return parMatch
	}

	for i := 0; i+wlen <= ls; i++ {
		match := make([]int, lw)
		j := 0
		for ; j < wlen; j++ {
			r := find(s[j+i],match)
			//fmt.Println(j + i, r)
			if !r {
				break
			}
		}
		if j == wlen {
			ans = append(ans, i)
		}
	}
	return ans
}
func main() {
	s := "ababaab"
	words := []string{"ab", "ba","ba"}

	fmt.Printf("%v\n", findSubstring(s, words))
	s = "bccbcc"
	words = []string{"bc","cc","cb"}

	fmt.Printf("%v\n", findSubstring(s, words))

	s = "barfoothefoobarman"
	words = []string{"foo","bar"}
	fmt.Printf("%v\n", findSubstring(s, words))
	s = "aaa"
	words = []string{"a","a"}
	fmt.Printf("%v\n", findSubstring(s, words))
	s = "abaababbaba"
	words = []string{"ab","ba","ab","ba"}
	fmt.Printf("%v\n", findSubstring(s, words))

	m := make(map[int]int)

	m[1] = 1
	m[2] = 2
	for k := range m{
		m[k] += 1
		fmt.Println(m[k])
	}

}
