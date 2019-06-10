package main

import (
	"fmt"
	"sort"
)


//nums1 = [1, 3]
//nums2 = [2]

//The median is 2.0

func main(){
	nums1 := []int{1,2,3,6}
	nums2 := []int{4,5,6,7}
	result := findMedOf2Arrs(nums1,nums2)
	fmt.Println(result)
}

//求两个有序数组的中位数

func findMedOf2Arrs(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1,nums2...)
	sort.Ints(nums1)
	nums1Len := len(nums1)
	var res float64
	if nums1Len %2 ==0{
		temp := nums1Len/2
		res = float64(nums1[temp-1]+nums1[temp])/2
		return res
	}
	//
	temp := nums1Len/2
	res = float64(nums1[temp])
	return res
}
