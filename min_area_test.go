package Algorithm

import (
	"fmt"
	"math"
	"testing"
)

func Test_MinArea(t *testing.T)  {
	nums := [][]int{
		{1,3,15,16,27,28},
		{2,6,8,14,17,38},
		{5,16,18,24,37,48},
	}
	ret := smallestRange(nums)
	fmt.Println(ret)
}

func smallestRange(nums [][]int) []int {
	var retTuple Tuple;
	for {
		if isAllHasElement(nums) {
			minIndex, min := getMinAtIndex0(nums)

			tuple := getMinArea(min, nums)
			if retTuple.empty() || tuple.length() < retTuple.length() {
				retTuple = tuple
			}
			nums[minIndex] = nums[minIndex][1:]
		} else {
			break
		}
	}

	fmt.Print(retTuple)
	return []int{retTuple.begin, retTuple.end}

}

func isAllHasElement(nums [][]int) bool {
	for _, arr := range nums {
		if len(arr) <= 0 {
			return false
		}
	}
	return true
}

func getMinArea(min int, nums [][]int) Tuple {
	var distance int;
	for _, arr := range nums {
		if arr[0] > distance {
			distance = arr[0]
		}
	}
	return Tuple{begin: min, end: distance}
}

func getMinAtIndex0 (nums [][]int) (int, int) {
	min := math.MaxInt64
	minIndex := -1
	for index, arr := range nums {
		if arr[0] < min {
			min = arr[0]
			minIndex = index;
		}
	}
	return minIndex, min
}

type Tuple struct {
	begin 	int;
	end 	int;
}

func (t Tuple)length() int {
	return t.end - t.begin
}

func (t Tuple)empty() bool {
	return t.begin == 0 && t.end == 0
}

func (t Tuple)String() string {
	return fmt.Sprintf("[%d, %d]", t.begin, t.end)
}

