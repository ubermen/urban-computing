package ds

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestKVMinHeap(t *testing.T) {
	h := &KVMinHeap{KV{2.0, "a"}, KV{1.0, "b"}, KV{5.0, "c"}}
	heap.Init(h)
	heap.Push(h, KV{3.0, "d"})
	fmt.Printf("minimum: %f\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%f ", heap.Pop(h))
	}
}
