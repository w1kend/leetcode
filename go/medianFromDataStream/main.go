package main

import (
	"container/heap"
	"fmt"
	"sort"
	"time"
)

type MedianFinder struct {
	Max *maxHeap
	Min *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		Max: &maxHeap{&intHeap{}},
		Min: &minHeap{&intHeap{}},
	}
}

func (this *MedianFinder) AddNum(num int) {

	if this.Min.Len() == 0 {
		heap.Push(this.Min, num)
		return
	}

	if this.Min.Len() == this.Max.Len() {
		if num < this.Max.Peek() {
			heap.Push(this.Min, heap.Pop(this.Max))
			heap.Push(this.Max, num)
		} else {
			heap.Push(this.Min, num)
		}
	} else {
		if num > this.Min.Peek() {
			heap.Push(this.Max, heap.Pop(this.Min))
			heap.Push(this.Min, num)
		} else {
			heap.Push(this.Max, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Min.Len() == this.Max.Len() {
		return float64(this.Max.Peek()+this.Min.Peek()) / 2.0
	}

	return float64(this.Min.Peek())
}

func main() {
	obj := Constructor()

	start := time.Now().Nanosecond()
	nums := []int{3, 2, 5, 7, 3, 4, 2, 3, 6, 8, 6, 5, 5, 10, 7}
	for _, v := range nums {
		obj.AddNum(v)
	}
	fmt.Println(obj.FindMedian())
	sort.Ints(nums)
	fmt.Printf("%+v, %d\n", nums, len(nums))
	//
	//
	//for i:= 0; i < 100000; i++ {
	//	obj.AddNum(rand.Int())
	//	obj.FindMedian()
	//}

	end := time.Now().Nanosecond()

	fmt.Printf("time - %d\n", (end - start))

	fmt.Println("min", obj.Min.intHeap, " ", obj.Min.Len())
	fmt.Println("max", obj.Max.intHeap, " ", obj.Max.Len())
	fmt.Println("========")
	l := obj.Min.Len()
	for i := 0; i < l; i++ {
		fmt.Print(heap.Pop(obj.Min), ", ")
	}
	fmt.Println("")
	l = obj.Max.Len()
	for i := 0; i < l; i++ {
		fmt.Print(heap.Pop(obj.Max), ", ")
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type intHeap []int

func (h intHeap) Len() int      { return len(h) }
func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h intHeap) Peek() int {
	return h[0]
}

type minHeap struct {
	*intHeap
}

func (h minHeap) Less(i, j int) bool { return (*h.intHeap)[i] < (*h.intHeap)[j] }

type maxHeap struct {
	*intHeap
}

func (h maxHeap) Less(i, j int) bool { return (*h.intHeap)[i] > (*h.intHeap)[j] }
