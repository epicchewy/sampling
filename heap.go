package sampling

type Node struct {
	Key float64
	Obj interface{}
}

// MinHeap implements a min heap based on Node.Key
// For weighted reservoir sampling, the Key is a relative weight to determine whether
// or not should be selected into the sample.
type MinHeap []*Node

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Key > h[j].Key
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *MinHeap) Pop() interface{} {
	curr := *h
	size := len(curr)
	*h = curr[0 : size-1]
	min := curr[size-1]
	return min
}

func (h *MinHeap) Peek() interface{} {
	return (*h)[len(*h)-1]
}
