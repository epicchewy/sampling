package sampling_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/epicchewy/sampling"
	"github.com/stretchr/testify/assert"
)

// Very naive tests to see if results are random samples
func TestSimpleReservoir(t *testing.T) {
	var items []int
	for i := 0; i < 100; i++ {
		items = append(items, i)
	}

	itr := sampling.NewIntInterator(items)
	r := sampling.NewSimpleReservoir(10, itr)
	samples := r.Sample()

	itr2 := sampling.NewIntInterator(items)
	r2 := sampling.NewSimpleReservoir(10, itr2)
	samples2 := r2.Sample()

	assert.NotEqual(t, samples, samples2)
}

func TestReservoir(t *testing.T) {
	var items []int
	for i := 0; i < 100; i++ {
		items = append(items, i)
	}

	itr := sampling.NewIntInterator(items)
	r := sampling.NewReservoir(10, itr)
	samples := r.Sample()

	itr2 := sampling.NewIntInterator(items)
	r2 := sampling.NewReservoir(10, itr2)
	samples2 := r2.Sample()

	assert.NotEqual(t, samples, samples2)
}

func TestWeightedReservoir(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var items []sampling.WeightedItem
	for i := 0; i < 100; i++ {
		items = append(
			items,
			sampling.WeightedItem{
				Weight: rand.Float64(),
				Obj:    i,
			},
		)
	}

	itr := sampling.NewItemIterator(items)
	r := sampling.NewWeightedReservoir(10, itr)
	samples := r.Sample()

	itr2 := sampling.NewItemIterator(items)
	r2 := sampling.NewWeightedReservoir(10, itr2)
	samples2 := r2.Sample()

	assert.NotEqual(t, samples, samples2)
}
