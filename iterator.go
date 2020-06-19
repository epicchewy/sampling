package sampling

import "bufio"

type Iterator interface {
	Next() (interface{}, bool)
}

type IntInterator struct {
	idx   int
	items []int
}

func NewIntInterator(items []int) Iterator {
	return &IntInterator{
		idx:   0,
		items: items,
	}
}

func (ii *IntInterator) Next() (interface{}, bool) {
	if ii.idx < len(ii.items) {
		item := ii.items[ii.idx]
		ii.idx++
		return item, true
	}
	return nil, false
}

type StringInterator struct {
	idx   int
	items []string
}

func NewStringInterator(items []string) Iterator {
	return &StringInterator{
		idx:   0,
		items: items,
	}
}

func (ii *StringInterator) Next() (interface{}, bool) {
	if ii.idx < len(ii.items) {
		item := ii.items[ii.idx]
		ii.idx++
		return item, true
	}
	return nil, false
}

type ChannelIterator struct {
	src chan interface{}
}

func NewChannelIterator(src chan interface{}) Iterator {
	return ChannelIterator{
		src: src,
	}
}

func (ci ChannelIterator) Next() (interface{}, bool) {
	item, ok := <-ci.src
	if !ok {
		return nil, false
	}
	return item, true
}

type BufferedReaderIterator struct {
	r *bufio.Reader
}

func NewBufferedReaderIterator(reader *bufio.Reader) Iterator {
	return BufferedReaderIterator{
		r: reader,
	}
}

func (bri BufferedReaderIterator) Next() (interface{}, bool) {
	item, err := bri.r.ReadString(byte('\n'))
	if err != nil {
		return nil, false
	}
	return item[:len(item)-1], true
}

type WeightedItem struct {
	Weight float64
	Obj    interface{}
}

type ItemIterator struct {
	idx   int
	items []WeightedItem
}

func NewItemIterator(items []WeightedItem) Iterator {
	return &ItemIterator{
		idx:   0,
		items: items,
	}
}

func (ii *ItemIterator) Next() (interface{}, bool) {
	if ii.idx < len(ii.items) {
		item := ii.items[ii.idx]
		ii.idx++
		return item, true
	}
	return nil, false
}
