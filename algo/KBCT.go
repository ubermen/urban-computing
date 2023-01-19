package algo

import (
	"container/heap"
	"github.com/tidwall/rtree"
	"math"
	"urban-computing-examples/ds"
	"urban-computing-examples/util"
)

func KBCT(k int, Q []util.Point, T []util.Line) []util.Line {
	if len(T) < k {
		println("k is bigger than number of trajectories")
		return nil
	}

	// λ-NN search를 위한 R-tree 구성
	var tr rtree.RTree
	for ri, R := range T {
		for _, r := range R.Points {
			point := [2]float64{r.X, r.Y}
			tr.Insert(point, point, ri)
		}
	}

	lambda := k
	delta := 1
	for true {
		// C 구성
		candidateWithLb := make(map[int]float64)
		ub := 0.0
		for _, q := range Q {
			matched := make(map[int]bool)
			point := [2]float64{q.X, q.Y}
			count := lambda
			tr.Nearby(
				rtree.BoxDist[float64, any](point, point, nil),
				func(min [2]float64, max [2]float64, data interface{}, dist float64) bool {
					ri := data.(int)
					// 단일 쿼리 포인트에서 특정 경로에 λ-NN이 여러개 있을 경우 가장 가까운 것을 사용한다.
					if !matched[ri] {
						matched[ri] = true
						candidateWithLb[ri] += math.Pow(1/math.E, dist)
					}
					count--
					if count == 0 {
						ub += math.Pow(1/math.E, dist)
						return false
					}
					return true
				},
			)
		}

		if len(candidateWithLb) >= k {
			// check if kth maximal lb of C >= ub of rest
			kthMaxLb := kthMaxLowerbound(k, candidateWithLb)
			if kthMaxLb >= ub {
				return refine(k, Q, T, candidateWithLb)
			}
		}
		lambda += delta
	}
	return nil
}

// TODO: ub-descending list 구성하고, k-BCT.min >= current.ub 이용해서 검색 시간 줄이기 최적화
func refine(k int, Q []util.Point, T []util.Line, candidateWithLb map[int]float64) []util.Line {
	h := &ds.KVMinHeap{}
	heap.Init(h)
	for ri, _ := range candidateWithLb {
		sim := similarity(Q, T[ri])
		kv := ds.KV{sim, T[ri]}
		if h.Len() < k {
			h.Push(kv)
		} else if (*h)[0].Key < sim {
			(*h)[0] = kv
			heap.Fix(h, 0)
		}
	}

	heap.Fix(h, 0)
	var kbct []util.Line
	for h.Len() > 0 {
		kbct = append(kbct, util.Line{heap.Pop(h).(ds.KV).Value.(util.Line).Points})
	}
	return kbct
}

func similarity(Q []util.Point, R util.Line) float64 {
	sum := 0.0
	closestDist := closestDist(Q, R)
	for _, dist := range closestDist {
		sum += math.Pow(1/math.E, dist)
	}
	return sum
}

func kthMaxLowerbound(k int, candidateWithLb map[int]float64) float64 {
	h := &ds.Float64MinHeap{}
	heap.Init(h)
	for _, lb := range candidateWithLb {
		if h.Len() < k {
			h.Push(lb)
		} else if (*h)[0] < lb {
			(*h)[0] = lb
			heap.Fix(h, 0)
		}
	}
	heap.Fix(h, 0)
	return (*h)[0]
}

func closestDist(Q []util.Point, R util.Line) []float64 {
	result := []float64{}
	for _, q := range Q {
		minDist := math.MaxFloat64
		for _, r := range R.Points {
			dist := util.PointDist(q, r)
			if dist < minDist {
				minDist = dist
			}
		}
		result = append(result, minDist)
	}
	return result
}
