package ds

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestFloat64MinHeap(t *testing.T) {
	h := &Float64MinHeap{2.0, 1.0, 5.0}
	heap.Init(h)
	heap.Push(h, 3.0)
	fmt.Printf("minimum: %f\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%f ", heap.Pop(h))
	}
}
