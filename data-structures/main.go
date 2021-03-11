package main

import (
	"fmt"
)

func main() {
	key1 := "key1"

	value1 := "val1"
	time1 := 123

	value2 := "val2"
	time2 := 234

	vMap := NewVersionedMap()
	vMap.Put(key1, value1, int64(time1))
	vMap.Put(key1, value2, int64(time2))

	//value, err := vMap.Get(key1, int64(1))
	//value, err := vMap.Get(key1, int64(125))
	value, err := vMap.Get(key1, int64(235))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}
