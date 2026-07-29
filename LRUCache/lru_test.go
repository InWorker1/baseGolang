package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("Последовательное вытеснение", func(t *testing.T) {
		cache := NewCacheStruct[string, int](2)
		cache.Set("f", 1)
		cache.Set("s", 2)
		cache.Set("fr", 3)
		if _, ok := cache.Get("f"); ok {
			t.Error()
		}
		cache.Set("s", 2)
		if val, _ := cache.buf["fr"]; val != cache.tail { // s стал первым
			t.Error()
		}
		if val, _ := cache.Get("fr"); cache.head.val != val {
			t.Error()
		}
	})

	t.Run("Отчистка кэша", func(t *testing.T) {
		cache := NewCacheStruct[int, int](4)
		cache.Set(1, 1)
		cache.Set(2, 2)
		cache.Set(3, 3)
		cache.Set(4, 4)
		if (cache.head.val != cache.buf[4].val) || (cache.tail.val != cache.buf[1].val) {
			t.Error("sets")
		}
		cache.Clear()
		if len(cache.buf) != 0 {
			t.Error("clear()")
		}
	})
}

// func TestLru(t *testing.T) {
// 	type step struct {
// 		funcName     string
// 		key          string
// 		val          int
// 		expectedVal  int
// 		expectedBool bool
// 	}
// 	tests := []struct {
// 		name     string
// 		capacity int
// 		steps    []step
// 	}{{name: "Последовательное вытеснение",
// 		capacity: 2, steps: {
// 			{funcName: "Set", key: "a", val: 10},
// 			{funcName: "Set", key: "b", val: 10},
// 			{funcName: "Set", key: "c", val: 10},
// 			{funcName: "Get", key: "a", val: 10},
// 		}}}
// }
