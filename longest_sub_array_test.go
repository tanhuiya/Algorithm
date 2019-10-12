package Algorithm

import (
	"fmt"
	"testing"
)

func Test_Longest_SubArray(t *testing.T)  {
	A := []int{0,1,1,1,1}
	B := []int{1,0,1,0,1}
	res := findLength(A, B)
	fmt.Println(res)
}

func findLength(A []int, B []int) int {
	dp := make([][]int, len(A)+ 1)
	for index, _ := range dp{
		row := make([]int, len(B) + 1)
		dp[index] = row
	}
	max := 0
	for i := 1; i < len(A) + 1; i++ {
		for j := 1; j < len(B) + 1; j++ {
			if A[i - 1] == B[j - 1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] == 3{
					fmt.Println("aa")
				}
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	fmt.Println(dp)
	return max
}
