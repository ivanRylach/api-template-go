package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionedMap_Get(t *testing.T) {

	key1 := "key1"

	value1 := "val1"
	time1 := 123

	value2 := "val2"
	time2 := 234

	vMap := NewVersionedMap()
	vMap.Put(key1, value1, int64(time1))
	vMap.Put(key1, value2, int64(time2))

	value, err := vMap.Get(key1, int64(1))
	assert.Error(t, err)

	value, err = vMap.Get(key1, int64(125))
	assert.Equal(t, value1, value)
	assert.NoError(t, err)

	value, err = vMap.Get(key1, int64(235))
	assert.Equal(t, value2, value)
	assert.NoError(t, err)

}

