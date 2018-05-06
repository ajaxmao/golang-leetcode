/*
 * [295] Find Median from Data Stream
 *
 * https://leetcode.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (29.77%)
 * Total Accepted:    60.7K
 * Total Submissions: 203.6K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n[[],[1],[2],[],[3],[]]'
 *
 * Median is the middle value in an ordered integer list. If the size of the
 * list is even, there is no middle value. So the median is the mean of the two
 * middle value.
 * Examples:
 * [2,3,4] , the median is 3
 * [2,3], the median is (2 + 3) / 2 = 2.5
 *
 *
 * Design a data structure that supports the following two operations:
 *
 *
 * void addNum(int num) - Add a integer number from the data stream to the data
 * structure.
 * double findMedian() - Return the median of all elements so far.
 *
 *
 *
 * For example:
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 *
 * Credits:Special thanks to @Louis1992 for adding this problem and creating
 * all test cases.
 */
import "fmt"

type MedianFinder struct {
	Nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	fmt.Println("add num", num)
	l := len(this.Nums)
	if l == 0 {
		this.Nums = []int{num}
		return
	}

	hi := l - 1
	var lo, mid int

	for lo <= hi {
		//小技巧：在这里不用 mid = (low + high)/2， 而是用 mid=low+((high-low)/2)，原因是 使用(low+high)/2会有整数溢出的问题。问题会出现在当low+high的结果大于表达式结果类型所能表示的最大值时，这样，产生溢出后再/2不会产生正确结果
		mid = lo + (hi-lo)/2
		midV := this.Nums[mid]
		if num == midV {
			break
		} else if num < midV {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if this.Nums[mid] <= num {
		mid++
	}
	fmt.Println("index is", mid)
	if mid == -1 {
		this.Nums = append([]int{num}, this.Nums...)
		return
	}
	left := append([]int{}, this.Nums[mid:l]...)

	this.Nums = append(append(this.Nums[:mid], num), left...)
}

func (this *MedianFinder) FindMedian() float64 {
	fmt.Println("find", this.Nums)
	ln := len(this.Nums)
	if ln == 0 {
		return 0.0
	}
	if ln%2 == 1 {
		return float64(this.Nums[ln/2])
	} else {
		return float64(this.Nums[ln/2]+this.Nums[ln/2-1]) / 2.0
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
