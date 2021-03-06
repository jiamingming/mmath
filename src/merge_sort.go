package main

import "fmt"

func main() {
	arr := []int{2,2,3,4,5,1,2,9,7,3,6,7,8,9,3,4,9,5,2,1}
	result := mergeSort(arr)
	fmt.Println(result)
}

/**
归并排序（MERGE-SORT）是建立在归并操作上的一种有效的排序算法,该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
时间复杂度为：O(n*log(n))
java 中Arrays.sort() ，就是用归并排序来实现的.

*/

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 // 
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...) 
	result = append(result, left[m:]...)
	return result
}





func bubbleSort(arr []int) {



}
