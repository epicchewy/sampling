package sampling

import (
	"container/heap"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Reservoir represents any of the following reservoirs (simple, normal, weighted)
// Each reservoir makes the assumption that the iterator contains more items than the desired samples
// and that the iterator is not an infinite stream of data.
type Reservoir interface {
	Sample() []interface{}
}

type simpleReservoir struct {
	k int
	i Iterator
}

// NewSimpleReservoir ...
func NewSimpleReservoir(samples int, itr Iterator) Reservoir {
	return simpleReservoir{
		k: samples,
		i: itr,
	}
}

func (sr simpleReservoir) Sample() []interface{} {
	var samples []interface{}

	// load samples with first k elements from iterator
	for i := 0; i < sr.k; i++ {
		item, hasNext := sr.i.Next()
		if !hasNext {
			panic("iterator must have more items than desired samples")
		}
		samples = append(samples, item)
	}

	n := sr.k

	for {
		item, hasNext := sr.i.Next()
		if !hasNext {
			break
		}
		n++
		j := rand.Intn(n)
		if j < sr.k {
			samples[j] = item
		}
	}

	return samples
}

type reservoir struct {
	k int
	// assume that super set of items can fit in memory
	i []interface{}
}

// NewReservoir ...
func NewReservoir(samples int, itr Iterator) Reservoir {
	var items []interface{}

	for {
		item, hasNext := itr.Next()
		if !hasNext {
			break
		}
		items = append(items, item)
	}

	return reservoir{
		k: samples,
		i: items,
	}
}

func (r reservoir) Sample() []interface{} {
	var samples []interface{}

	// load samples with first k elements from iterator
	for i := 0; i < r.k; i++ {
		samples = append(samples, r.i[i])
	}

	w := math.Pow(math.E, math.Log(rand.Float64())/float64(r.k))
	i := float64(0)
	n := float64(len(r.i))

	for i < n {
		i = i + math.Floor(math.Log(rand.Float64())/math.Log(1-w)) + 1.0
		if i < n {
			samples[rand.Intn(r.k)] = r.i[int(i)]
			w = w * math.Pow(math.E, math.Log(rand.Float64())/float64(r.k))
		}
	}

	return samples
}

type weightedReservoir struct {
	k int
	i Iterator
}

// NewWeightedReservoir ...
func NewWeightedReservoir(samples int, itr Iterator) Reservoir {
	return weightedReservoir{
		k: samples,
		i: itr,
	}
}

func (wr weightedReservoir) Sample() []interface{} {
	var samples []interface{}

	h := &MinHeap{}
	heap.Init(h)

	for {
		item, hasNext := wr.i.Next()
		if !hasNext {
			break
		}

		wi, ok := item.(WeightedItem)
		if !ok {
			panic("weighted reservoir can only accept items of type WeightedItem")
		}

		r := math.Pow(rand.Float64(), 1/wi.Weight)

		if h.Len() < wr.k {
			heap.Push(h, &Node{Key: r, Obj: wi.Obj})
		} else {
			head := h.Peek().(*Node)
			if r > head.Key {
				_ = heap.Pop(h)
				heap.Push(h, &Node{Key: r, Obj: wi.Obj})
			}
		}
	}

	for h.Len() > 0 {
		samples = append(samples, heap.Pop(h).(*Node).Obj)
	}

	return samples
}
