package structs

type Queue []int

func (h Queue) Len() int           { return len(h) }
func (h Queue) Less(i, j int) bool { return true }
func (h Queue) Swap(i, j int)      { return }

func (h *Queue) Push(x interface{}) {
	// Add and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Queue) Pop() interface{} {
	old := *h
	n := len(old)
	if n == 0 {
		panic("failed to Pop() from empty queue")
	}
	x := old[0]
	*h = old[1:n]
	return x
}
