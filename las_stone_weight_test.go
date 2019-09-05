package Algorithm

import (
	"fmt"
	"testing"
)

func Test_last_stone(t *testing.T)  {
	nums := []int{2,7,4,1,8,1,40}
	ret := lastStoneWeightII(nums)
	fmt.Println(ret)
}

func lastStoneWeightII(stones []int) int {
	total := 0
	for _, n := range stones {
		total += n
	}

	sum := total / 2
	dp := make([]int, sum + 1)
	for _, stone := range stones {
		for j := sum; j - stone >=0 ; j-- {
			dp[j] = max2(dp[j], dp[j - stone] + stone)
		}
	}
	return total - 2 * dp[sum]
}

func max2(i , j int) int {
	if i > j {
		return i
	}
	return j
}