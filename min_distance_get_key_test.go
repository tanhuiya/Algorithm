package Algorithm

import (
	"fmt"
	"testing"
)

func TestGetKey(t *testing.T) {
	grid := []string{
		"@...a",
		".###A",
		"b.BCc",
	}
	length := shortestPathAllKeys(grid)
	fmt.Println(length)
}

func shortestPathAllKeys(grid []string) int {
	stateNum := getStateNum(grid)
	x, y := getStart(grid)
	min := minStep(grid, x, y, stateNum)
	return min
}

type Node struct {
	x int
	y int
	step int
	keyState int
	pre *Node
}

func minStep(grid []string, row, col int, stateNum int) int {
	dir := [][2]int {
		{0, 1}, {1, 0},
		{0, -1}, {-1, 0},
	};
	step := 0
	var queue []Node
	queue = append(queue, Node{x: row, y: col, step: step, keyState: 0, pre: nil})
	visited := make([][][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([][]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = make([]bool, stateNum + 1)
		}
	}
	visited[row][col][0] = true
	for {
		if len(queue) < 1 {
			break
		}

		node := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			next := Node{
				x: node.x + dir[i][0],
				y: node.y + dir[i][1],
				step: node.step + 1,
				keyState: node.keyState,
				pre:  &node,
			}
			if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
				continue
			}
			val := grid[next.x][next.y]
			if val == '#' || visited[next.x][next.y][next.keyState] {
				continue
			}
			if val <= 'Z' && val >= 'A'{
				if !hasKeyforLock(val, next.keyState){
					continue
				}
			}
			if val <= 'z' && val >= 'a' {
				next.keyState |= 1 << (val - 'a')
				if next.keyState == stateNum {
					//for {
					//	if next.pre != nil {
					//		fmt.Println(next.x, next.y)
					//		next = *next.pre
					//	}
					//}
					return next.step
				}
			}
			visited[next.x][next.y][next.keyState] = true
			queue = append(queue, next)
		}
	}
	return -1
}

func hasKeyforLock(lock byte, keyState int) bool {
	return (1 << (lock - 'A') & keyState) > 0
}

func getStateNum(grid []string) int {
	stateNum := 0
	for _, row := range grid {
		for _, b := range row {
			if b >= 'a' && b <= 'z' {
				stateNum |= 1 << uint32(b - 'a')
			}
		}
	}
	return stateNum
}

func getStart(grid []string) (int, int) {
	for x, row := range grid {
		for y, b := range row {
			if b == '@' {
				return x,y
			}
		}
	}
	return 0, 0
}



