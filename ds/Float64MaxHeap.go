package ds

type Float64MaxHeap []float64

func (h Float64MaxHeap) Len() int           { return len(h) }
func (h Float64MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h Float64MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Float64MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(float64))
}

func (h *Float64MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
