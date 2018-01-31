/*
* 一维模式识别算法一，穷举算法
* 时间复杂度：　O(n³)
 */
package main

import "fmt"

func maxSoFar(a []int) int {
	var maxSoFar int
	len := len(a)
	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			var sum int
			for k := i; k < j; k++ {
				sum += a[k]
			}
			if sum > maxSoFar {
				maxSoFar = sum
			}
		}
	}
	return maxSoFar
}

func main() {
	a := []int{31, -42, 59, 26, -53, 58, 97, -93, -23, 84}
	max := maxSoFar(a)
	fmt.Println("max", max)
}
