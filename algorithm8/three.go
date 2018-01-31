/*
* 一维模式识别算法三，穷举法的优化算法，使用了预处理的方式
* 时间负责度： O(n²)
 */
package main

import (
	"fmt"
)

func maxSoFar(a []int) int {
	var maxSoFar int
	len := len(a)

	cumarr := make([]int, len)
	for i := 0; i < len; i++ {
		if i == 0 {
			cumarr[i] = a[i]
		} else {
			cumarr[i] = cumarr[i-1] + a[i]
		}
	}

	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			var sum int
			if i == 0 {
				sum = a[i]
			} else {
				sum = cumarr[j] - cumarr[i-1]
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
