package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T){
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T){
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("potatoes"),
		},
	}

	for _, cas := range cases{
		cache.Add(cas.inputKey, cas.inputVal)

		actual, ok := cache.Get(cas.inputKey)
		if !ok{
			t.Errorf("%v entry not found", cas.inputKey)
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("value in cache does not match %v vs %v", string(actual), string(cas.inputVal))
		}
	}
}

func TestREAP(t *testing.T){
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok{
		t.Errorf("cache still persists when it should have been deleted")
	}
}