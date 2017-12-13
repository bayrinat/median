package structs

// An MaxHeap is a max-heap of ints.
// Based on https://golang.org/src/container/heap/example_intheap_test.go
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Add and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) IndexOf(x interface{}) int {
	for i := 0; i < h.Len(); i++ {
		if x == (*h)[i] {
			return i
		}
	}
	return -1
}
