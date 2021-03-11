package main

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {

	items := []*Item{
		{
			value:    "banana",
			priority: 3,
			index:    0,
		},
		{
			value:    "apple",
			priority: 2,
			index:    1,
		},
		{
			value:    "pear",
			priority: 4,
			index:    2,
		},
	}

	pq := make(PriorityQueue, len(items))
	for i, item := range items {
		pq[i] = item
	}

	heap.Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}

	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	items = []*Item{}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		items = append(items, item)
	}

	expectedItems := []*Item{
		{
			value:    "orange",
			priority: 5,
			index:    -1,
		},
		{
			value:    "pear",
			priority: 4,
			index:    -1,
		},
		{
			value:    "banana",
			priority: 3,
			index:    -1,
		},
		{
			value:    "apple",
			priority: 2,
			index:    -1,
		},
	}
	//expectedItems := []*Item{
	//	{
	//		value:    "apple",
	//		priority: 2,
	//		index:    -1,
	//	},
	//	{
	//		value:    "banana",
	//		priority: 3,
	//		index:    -1,
	//	},
	//	{
	//		value:    "pear",
	//		priority: 4,
	//		index:    -1,
	//	},
	//	{
	//		value:    "orange",
	//		priority: 5,
	//		index:    -1,
	//	},
	//
	//}

	assert.True(t, reflect.DeepEqual(expectedItems, items))

}
