package Algorithm

import (
	"fmt"
	"testing"
)

func Test_Splite_Array(t *testing.T)  {
	nums := []int{7,2,5,10,8}
	m := 2
	res := splitArray(nums, m)
	fmt.Println(res)
}

func splitArray(nums []int, m int) int {
	max := -1
	sum := 0
	for _, val := range nums {
		if val > max {
			max = val
		}
		sum += val
	}
	l, r := max, sum
	for {
		mid := (l + r) / 2
		cnt, temp := 1, 0
		for _, val := range nums {
			if temp + val <= mid {
				temp += val
			} else {
				cnt++
				temp = val
			}
		}
		if cnt > m {
			l = mid + 1
		} else {
			r = mid
		}
		if l == r {
			return l
		}
	}
	return 0
}
