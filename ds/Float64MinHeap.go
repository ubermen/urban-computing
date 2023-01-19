package ds

type Float64MinHeap []float64

func (h Float64MinHeap) Len() int           { return len(h) }
func (h Float64MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h Float64MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Float64MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(float64))
}

func (h *Float64MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
