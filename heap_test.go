package sampling_test

import (
	"container/heap"
	"testing"

	"github.com/epicchewy/sampling"
	"github.com/stretchr/testify/assert"
)

func TestHeapPeek(t *testing.T) {
	h := &sampling.MinHeap{
		&sampling.Node{Key: 0.3, Obj: 3},
		&sampling.Node{Key: 0.6, Obj: 4},
		&sampling.Node{Key: 0.9, Obj: 5},
	}
	heap.Init(h)

	heap.Push(h, &sampling.Node{Key: 0.1, Obj: 1})
	top := h.Peek().(*sampling.Node)
	assert.Equal(t, top.Key, 0.1)
}

func TestHeapPop(t *testing.T) {
	h := &sampling.MinHeap{
		&sampling.Node{Key: 0.3, Obj: 3},
		&sampling.Node{Key: 0.6, Obj: 4},
		&sampling.Node{Key: 0.9, Obj: 5},
	}
	heap.Init(h)

	pop := h.Pop().(*sampling.Node)
	assert.Equal(t, pop.Key, 0.3)
}
