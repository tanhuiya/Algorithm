package time_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func test()  {
	//a := GetIndex([]int{2, 4, 6, 8}, 5)
	//assert.Equal(t, a, 1)
	//
	//a = GetIndex([]int{2, 4, 6, 8}, 10)
	//assert.Equal(t, a, 3)
	//
	//a = GetIndex([]int{}, 1)
	//assert.Equal(t, a, -1)
	//
	//a = GetIndex([]int{2, 4, 6, 8}, 1)
	//assert.Equal(t, a, -1)
	//
	//a = GetIndex([]int{2, 4, 6, 8}, 3)
	//assert.Equal(t, a, 0)
	//
	//a = GetIndex([]int{2, 4, 6, 8}, 7)
	//assert.Equal(t, a, 2)
}

func TestTimeMap(t *testing.T)  {
	obj := Constructor();
	//obj.Set("foo", "bar2", 2);
	//obj.Set("foo", "bar9", 9);
	//obj.Set("foo", "bar4", 4);
	//obj.Set("foo", "bar8", 8);
	//obj.Set("foo", "bar1", 1);
	//
	//param_3 := obj.Get("foo", 10);
	//assert.Equal(t, "bar9", param_3)
	//
	//param_3 = obj.Get("foo", 8);
	//assert.Equal(t, "bar8", param_3)
	//
	//param_3 = obj.Get("foo", 5);
	//assert.Equal(t, "bar4", param_3)
	//
	//param_3 = obj.Get("foo1", 5);
	//assert.Equal(t, "", param_3)

	obj.Set("foo", "bar", 1);
	param_3 := obj.Get("foo", 1);
	assert.Equal(t, "bar", param_3)
	param_3 = obj.Get("foo", 3);
	assert.Equal(t, "bar", param_3)
	obj.Set("foo", "bar2", 4);
	param_3 = obj.Get("foo", 4);
	assert.Equal(t, "bar2", param_3)
	param_3 = obj.Get("foo", 5);
	assert.Equal(t, "bar2", param_3)

}

type Pair struct {
	value 		string
	timestamp 	int
}

type TimeMap struct {
	tMap map[string][]Pair
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		tMap: make(map[string][]Pair),
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	if _, ok := this.tMap[key]; !ok {
		this.tMap[key] = []Pair{}
	}
	this.tMap[key] = insertPair(this.tMap[key], Pair{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	pairs := this.tMap[key]
	if len(pairs) == 0 {
		return ""
	}
	index := get(pairs, timestamp)
	if index == -1 {
		return ""
	}
	return pairs[index].value
}

// [1, 3, 5, 7, 9]
// 5
func get(pairs []Pair, timestamp int) int {
	start, end := 0, len(pairs) - 1
	mid := (start + end + 1) / 2

	if len(pairs) == 0 || pairs[0].timestamp > timestamp {
		return -1
	}
	for {
		if end <= start {
			break
		}
		if pairs[mid].timestamp > timestamp {
			end = mid - 1
		} else {
			start = mid
		}
		mid = (start + end + 1) / 2
	}

	return mid
}

func insertPair(pairs []Pair, item Pair) []Pair {
	idx := findIndex(pairs, item)
	if idx == -1 {
		return append([]Pair{item}, pairs...)
	}

	left := append([]Pair{}, pairs[idx + 1:]...)
	retpair := append(pairs[:idx + 1], item)
	retpair = append(retpair, left...)
	return retpair
}


func findIndex(pairs []Pair, item Pair) int {
	start, end := 0, len(pairs) - 1
	mid := (start + end + 1) / 2

	if len(pairs) == 0 || pairs[0].timestamp > item.timestamp {
		return -1
	}
	if item.timestamp > pairs[end].timestamp {
		return end
	}

	for {
		if start < end  {
			if pairs[mid].timestamp > item.timestamp {
				end = mid - 1
			} else {
				start = mid
			}
			mid = (start + end + 1) / 2
		} else {
			break
		}
	}

	return  mid
}
//
//func GetIndex(arr []int, target int) int {
//	start, end := 0, len(arr)  - 1
//	mid := (start + end + 1) / 2
//
//	if len(arr) == 0 || target < arr[0] {
//		return -1
//	}
//	if target > arr[end] {
//		return end
//	}
//
//	for {
//		if start < end {
//			if arr[mid] > target {
//				end = mid - 1
//			} else {
//				start = mid
//			}
//			mid = (start + end + 1) / 2
//		} else {
//			break
//		}
//	}
//	return mid
//}

