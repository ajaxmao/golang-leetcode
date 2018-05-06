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
import (
	"container/heap"
)

type MedianFinder struct {
	RightHeap intHeap
	LeftHeap  intHeap
	Total     int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	rHeap, lHeap := new(intHeap), new(intHeap)
	heap.Init(rHeap)
	heap.Init(lHeap)
	return MedianFinder{
		RightHeap: *rHeap,
		LeftHeap:  *lHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.Total++
	if this.Total%2 == 0 {
		heap.Push(&this.LeftHeap, -1*num)
	} else {
		heap.Push(&this.RightHeap, num)
	}
	if this.Total > 1 && this.RightHeap[0] < (-1*this.LeftHeap[0]) {
		this.RightHeap[0], this.LeftHeap[0] = -1*this.LeftHeap[0], -1*this.RightHeap[0]
		heap.Fix(&this.RightHeap, 0)
		heap.Fix(&this.LeftHeap, 0)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Total == 0 {
		return 0.0
	}
	if this.Total%2 == 0 {
		return float64(this.RightHeap[0]-this.LeftHeap[0]) / 2.0
	} else {
		return float64(this.RightHeap[0])
	}
}

// intHeap 实现了 heap 的接口
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
