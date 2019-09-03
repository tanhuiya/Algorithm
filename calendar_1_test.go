package Algorithm

import (
	"fmt"
	"testing"
)

func Test_Calendar(t *testing.T)  {
	obj := Constructor();
	param_1 := obj.Book(20,29);
	fmt.Println(param_1)
	param_1 = obj.Book(13,22);
	fmt.Println(param_1)

	param_1 = obj.Book(44,50);
	fmt.Println(param_1)

	param_1 = obj.Book(1,7);
	fmt.Println(param_1)


	param_1 = obj.Book(2,10);
	fmt.Println(param_1)

	param_1 = obj.Book(36,42);
	fmt.Println(param_1)
	//r,idx := book(obj.ranges, 30, 0)
}

type Range struct {
	start int
	end int
}

type MyCalendar struct {
	ranges []Range
}

func Constructor() MyCalendar {
	return MyCalendar{
		ranges: []Range{
		},
	}
}


func (this *MyCalendar) Book(start int, end int) bool {
	newRange := Range{start, end}
	if len(this.ranges) == 0 {
		this.ranges = append(this.ranges, Range{start, end})
		return true
	}

	ran, idx := book(this.ranges, start, 0)
	if idx == 0 {
		if newRange.end <= ran.start {
			temp := []Range{newRange}
			temp = append(temp, this.ranges...)
			this.ranges = temp
			return true
		} else if newRange.start >= ran.end {
			temp := append(this.ranges, newRange)
			this.ranges = temp
			return true
		} else {
			return false
		}
	} else if idx == len(this.ranges) - 1 &&
		newRange.start >= ran.end {
		temp := append(this.ranges, newRange)
		this.ranges = temp
		return true
	} else {
		pre := this.ranges[idx - 1]
		if newRange.start >= pre.end && newRange.end <= ran.start {
			//temp := append(this.ranges[:idx], newRange)
			temp := []Range{}
			temp = append(temp, this.ranges[:idx]...)
			temp = append(temp, newRange)
			temp = append(temp, this.ranges[idx:]...)
			this.ranges = temp
			return true
		} else {
			return false
		}
	}
	return false
}

func book(ranges []Range, start int, idx int) (Range, int) {
	if len(ranges) == 1 {
		return ranges[0], idx
	}
	middle := len(ranges) / 2 - 1
	midRange := ranges[middle]

	if start <= midRange.start {
		return book(ranges[:middle + 1], start, idx)
	} else {
		return book(ranges[middle + 1:], start, idx + middle + 1)
	}
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */