/*
* 一维模式识别算法五，扫描算法
* 时间复杂度：O(n)
 */
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSoFar(a []int) int {
	var maxSoFar int
	var maxEndingHere int
	len := len(a)

	for i := 0; i < len; i++ {
		maxEndingHere = max(maxEndingHere+a[i], 0)
		maxSoFar = max(maxEndingHere, maxSoFar)
	}
	return maxSoFar

}

func main() {
	a := []int{31, -42, 59, 26, -53, 58, 97, -93, -23, 84}
	max := maxSoFar(a)
	fmt.Println("max", max)

}
