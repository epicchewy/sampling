package sampling_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/epicchewy/sampling"
)

func TestIntIterator(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	itr := sampling.NewIntInterator(items)
	for _, expected := range items {
		item, hasNext := itr.Next()
		assert.Equal(t, expected, item)
		assert.True(t, hasNext)
	}
	item, hasNext := itr.Next()
	assert.Nil(t, item)
	assert.False(t, hasNext)
}

func TestStringIterator(t *testing.T) {
	items := []string{"foo", "bar", "baz"}
	itr := sampling.NewStringInterator(items)
	for _, expected := range items {
		item, hasNext := itr.Next()
		assert.Equal(t, expected, item)
		assert.True(t, hasNext)
	}
	item, hasNext := itr.Next()
	assert.Nil(t, item)
	assert.False(t, hasNext)
}

func TestChannelIterator(t *testing.T) {
	items := []interface{}{0, 1, "two", 3, "vier"}
	ch := make(chan interface{})
	go func() {
		for _, i := range items {
			ch <- i
		}
		close(ch)
	}()

	itr := sampling.NewChannelIterator(ch)
	for _, expected := range items {
		item, hasNext := itr.Next()
		assert.Equal(t, expected, item)
		assert.True(t, hasNext)
	}
	item, hasNext := itr.Next()
	assert.Nil(t, item)
	assert.False(t, hasNext)
}

func TestBufferedReaderIterator(t *testing.T) {
	data := "hello\nworld\n!\n"
	b := bufio.NewReader(strings.NewReader(data))
	itr := sampling.NewBufferedReaderIterator(b)
	expected := strings.Split(data, "\n")
	for i := 0; i < len(expected)-1; i++ {
		item, hasNext := itr.Next()
		assert.Equal(t, expected[i], item)
		assert.True(t, hasNext)
	}
	item, hasNext := itr.Next()
	assert.Nil(t, item)
	assert.False(t, hasNext)
}
