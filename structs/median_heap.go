package structs

import (
	"container/heap"
)

// MedianHeap wraps min and max heaps, the median are stored in median value
type MedianHeap struct {
	minHeap MinHeap
	maxHeap MaxHeap
	median  float64
}

// Returns new instance of MedianHeap
func NewMedianHeap() *MedianHeap {
	h := &MedianHeap{
		minHeap: MinHeap{},
		maxHeap: MaxHeap{},
		median:  0,
	}

	heap.Init(&h.minHeap)
	heap.Init(&h.maxHeap)

	return h
}

// Length of heap
func (h *MedianHeap) Len() int {
	return h.minHeap.Len() + h.maxHeap.Len()
}

// Adds the new value to collection. After that balance internal heaps and recalculate median
func (h *MedianHeap) Add(x int) {
	if float64(x) < h.median {
		heap.Push(&h.maxHeap, x)
	} else {
		heap.Push(&h.minHeap, x)
	}

	h.balance()
	h.setMedian()
}

// Removes the value from collection. After that balance internal heaps and recalculate median
func (h *MedianHeap) Remove(x int) {
	h.remove(x)
	h.balance()
	h.setMedian()
}

// Returns heap's median value
func (h *MedianHeap) Median() float64 {
	return h.median
}

// Balance two internal heaps
func (h *MedianHeap) balance() {
	if h.minHeap.Len()-h.maxHeap.Len() > 1 {
		heap.Push(&h.maxHeap, heap.Pop(&h.minHeap))
	}
	if h.maxHeap.Len()-h.minHeap.Len() > 1 {
		heap.Push(&h.minHeap, heap.Pop(&h.maxHeap))
	}
}

// Recalculates median and stores it to h.median
func (h *MedianHeap) setMedian() {
	if h.Len() == 0 {
		return
	}

	if h.minHeap.Len() == h.maxHeap.Len() {
		h.median = (float64(h.minHeap[0]) + float64(h.maxHeap[0])) / 2
	} else {
		if h.maxHeap.Len() > h.minHeap.Len() {
			h.median = float64(h.maxHeap[0])
		} else {
			h.median = float64(h.minHeap[0])
		}
	}
}

// Finds out needed heap, finds index and removes the value
// TODO (bayrinat): simplify it
func (h *MedianHeap) remove(x int) {
	var index int

	if float64(x) == h.median {
		if h.maxHeap.Len() < h.minHeap.Len() {
			index = h.minHeap.IndexOf(x)
			heap.Remove(&h.minHeap, index)
		} else {
			index = h.maxHeap.IndexOf(x)
			heap.Remove(&h.maxHeap, index)
		}
	} else {
		if float64(x) < h.median {
			index := h.maxHeap.IndexOf(x)
			heap.Remove(&h.maxHeap, index)
		} else {
			index := h.minHeap.IndexOf(x)
			heap.Remove(&h.minHeap, index)
		}
	}
}
