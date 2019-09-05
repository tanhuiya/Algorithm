package Algorithm

import (
	"fmt"
	"testing"
)

func TestOneZero(t *testing.T)  {
	strs := []string{
		"10", "0001", "111001", "1", "0",
	}
	num := findMaxForm(strs, 5, 3)
	fmt.Println(num)
}

func findMaxForm(strs []string, m int, n int) int {
	countArr := coutStrs(strs)
	dp := make([][][]int, len(countArr) + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m + 1)
		for j := 0; j < m + 1; j++ {
			dp[i][j] = make([]int, n + 1)
		}
	}

	for i := 1; i <= len(countArr); i++ {
		tup := countArr[i - 1]
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if tup.m <= j && tup.n <= k {
					dp[i][j][k] = max1(dp[i - 1][j - tup.m][k - tup.n] + 1, dp[i - 1][j][k])
				} else {
					dp[i][j][k] = dp[i - 1][j][k]
				}
			}
		}
	}

	return dp[len(countArr)][m][n]
}

func max1(i , j int) int {
	if i > j {
		return i
	}
	return j
}

type Tup struct {
	m int
	n int
}
func coutStrs(strArr []string) ([]Tup) {
	tups := []Tup{}
	for _, value := range strArr {
		numZero := 0
		numOne := 0
		for _, ch := range value {
			if ch == '0' {
				numZero++
			} else {
				numOne++
			}
		}
		tups = append(tups, Tup{numZero, numOne})
	}
	return tups
}