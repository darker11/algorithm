/*
* 一维模式识别算法二，穷举法的优化算法
* 时间负责度： O(n²)
 */
package main

import (
	"fmt"
)

func maxSoFar(a []int) int {
	var maxSoFar int
	len := len(a)
	for i := 0; i < len; i++ {
		sum := 0
		for j := i; j < len; j++ {
			sum += a[j]

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
