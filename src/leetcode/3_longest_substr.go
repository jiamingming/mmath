package main

import "fmt"

//
func main(){
	s := `aaabcddde`
	start,max := lenOfLongestSubstring(s)
	fmt.Println(start,max)
}


//返回符合条件的子串的起始index和length
func lenOfLongestSubstring(s string) (int,int) {
	m := make(map[rune]int)
	start, max := -1, 0

	for k, v := range s {
		//repeated
		if last, ok := m[v]; ok && last > start {
			start = last
		}
		m[v] = k
		//max
		if k-start > max {
			max = k - start
		}
	}
	sindex := start-max
	return sindex,max
}
