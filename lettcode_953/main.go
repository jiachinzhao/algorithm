package main


func isAlienSorted(words []string, order string) bool {
	rank := make(map[rune]int)
	for index, v := range order{
		rank[v] = index
	}
	cmp := func(s1, s2 string) bool {
		len1 := len(s1)
		len2 := len(s2)
		for i := 0;i < len1 && i < len2;i++{
			if rank[rune(s1[i])] < rank[rune(s2[i])] {
				return true
			} else if rank[rune(s1[i])] == rank[rune(s2[i])] {
				continue
			}
			return false
		}
		return len1 <= len2
	}
	for i := 0;i < len(words) - 1;i++{
		if !cmp(words[i], words[i + 1])	{
			return false
		}
	}
	return true
}
func main(){

}



