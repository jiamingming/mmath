package main

import "fmt"
func main(){

	arr := []int{3,5,9,8,7,4,3,7,6}
	result := HeapSort(arr)
	fmt.Println(result)
}


func minHeap(root int, end int, c []int)  {
   for {
      var child = 2*root + 1
      //child node 
      if child > end {
         break
      }
      //right chinld node
      if child+1 <= end && c[child] > c[child+1] {
         child += 1
      }
      if c[root] > c[child] {
         c[root], c[child] = c[child], c[root]
         root = child
      } else {
         break
      }
   }
}
//desc
func HeapSort(c []int) []int {
   var n = len(c)-1
   for root := n / 2; root >= 0; root-- {
      minHeap(root, n, c)
   }
   fmt.Println("")
   for end := n; end >=0; end-- {
      if c[0]<c[end]{
         c[0], c[end] = c[end], c[0]
         minHeap(0, end-1, c)
      }
   }
   return c
}
