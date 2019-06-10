package main

import "fmt"

func main(){
	nums := []int{1,2,5,6,8,9}
	target := 7
	result := twoSum(nums,target)
	fmt.Println(result)
}

//leetcode two_sum
func twoSum(nums []int, target int) []int {
    h := make(map[int]int)
    for i, value := range nums {
        if wanted, ok := h[value]; ok {
            return []int{wanted, i}
            //fmt.Println(wanted,i)
        } else {
            h[target-value] = i
        }
    }
    return nil
}
