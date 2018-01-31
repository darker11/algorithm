/*
* 一维模式识别算法四，分治算法
* 时间复杂度：O(nlogn)
 */
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func maxSoFar(a []int, l, r int) int {
	if l > r {
		return 0
	}
	if l == r {
		return max(0, a[l])
	}
	m := (l + r) / 2
	var lmax int
	var sum int
	var rmax int
	//此处必须从m往l移动（从中间往左边移动），因为跨越边界的最大向量由左边和右边两部分的和组成
	for i := m; i >= l; i-- {
		sum += a[i]
		lmax = max(lmax, sum)
	}

	sum = 0 //此处必须重置为０，否则会影响后面的计算

	//此处必须从m往r移动（从中间往右边移动），因为跨越边界的最大向量由左边和右边两部分的和组成
	for i := m + 1; i <= r; i++ {
		sum += a[i]
		rmax = max(rmax, sum)
	}
	fmt.Printf("lmax: %d, rmax: %d, m: %d\n", lmax, rmax, m)
	return max3(lmax+rmax, maxSoFar(a, l, m), maxSoFar(a, m+1, r))
}

func main() {
	a := []int{31, -42, 59, 26, -53, 58, 97, -93, -23, 84}
	max := maxSoFar(a, 0, len(a)-1)
	fmt.Println("max", max)

}
