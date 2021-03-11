package main

import (
	"errors"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
	"time"
)

type VersionedMap struct {
	versionedMap *hashmap.Map
}

func NewVersionedMap() *VersionedMap {
	return &VersionedMap{versionedMap: hashmap.New()}
}

func (vm *VersionedMap) Put(key, value string, unixSec int64) {
	iVersionedValue, found := vm.versionedMap.Get(key)
	if !found {
		iVersionedValue = treemap.NewWith(utils.TimeComparator)
		vm.versionedMap.Put(key, iVersionedValue)
	}
	versionedValue := iVersionedValue.(*treemap.Map)

	unixTime := time.Unix(unixSec, 0)

	versionedValue.Put(unixTime, value)
}

func (vm *VersionedMap) Get(key string, unixSec int64) (value string, err error) {
	iVersionedValue, found := vm.versionedMap.Get(key)
	if !found {
		return "", errors.New("key not found")
	}
	versionedValue := iVersionedValue.(*treemap.Map)
	if versionedValue.Empty() {
		return "", errors.New("no values found")
	}
	unixTime := time.Unix(unixSec, 0)

	_, foundValue := versionedValue.Floor(unixTime)

	if foundValue != nil {
		return foundValue.(string), nil
	} else {
		return "", errors.New("no values recorded prior to the requested time")
	}
}
