package Algorithm

import (
	"fmt"
	"testing"
)

func Test_Partition(t *testing.T) {
	nums := []int{1, 2, 5, 4}
	can := canPartition(nums)
	fmt.Println(can)
}

func canPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total % 2 == 1 {
		return false
	}
	total = total / 2
	w := make([]int, total + 1)

	dp := make([][]int, len(nums) + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(w))
	}
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= total; j++ {
			val := nums[i - 1]

			if j < val {
				dp[i][j] = dp[i - 1][j]
			} else {
				dp[i][j] = max(dp[i - 1][j - val] + val, dp[i - 1][j])
			}
			if dp[i][j] == total {
				return true
			}
		}
	}
	return false
}

func max(i , j int) int {
	if i > j {
		return i
	}
	return j
}