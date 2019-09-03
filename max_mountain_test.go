package Algorithm

import (
	"fmt"
	"testing"
)

func TestMountain(t *testing.T) {
	//arr := []int{2,1,4,7,3,2,5}
	arr := []int{2,4,3}
	length := longestMountain(arr)
	fmt.Println(length)
}

func longestMountain(arr []int) int {
	longest := 0
	for index, val := range arr {
		if index + 3 > len(arr) {
			break
		}
		next := arr[index + 1]

		if val >= next {
			continue
		}

		length := getMoutainForIndex(index, arr)
		if length > longest {
			longest = length
		}
	}
	return longest
}

func getMoutainForIndex(index int, arr []int) int {
	orgIndex := index
	reachTop := false
	for {
		// 到最后一个，还没到 top
		if  index == len(arr) - 1 {
			if !reachTop {
				return 0
			} else {
				return index - orgIndex + 1
			}
		}

		if !reachTop {
			if arr[index] == arr[index + 1] {
				return 0
			}
			if arr[index] > arr[index + 1] {
				reachTop = true
			}
			index = index + 1
		} else {
			if arr[index] > arr[index + 1] {
				index = index + 1
			} else {
				return index - orgIndex + 1
			}
		}

	}
}
