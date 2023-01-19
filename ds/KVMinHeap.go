package ds

type KV struct {
	Key   float64
	Value interface{}
}
type KVMinHeap []KV

func (h KVMinHeap) Len() int           { return len(h) }
func (h KVMinHeap) Less(i, j int) bool { return h[i].Key < h[j].Key }
func (h KVMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KVMinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(KV))
}

func (h *KVMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
