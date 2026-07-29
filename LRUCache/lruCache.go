package main

import "sync"

type Cache[K comparable, V any] struct {
	buf      map[K]*Node[K, V]
	capacity int
	head     *Node[K, V]
	tail     *Node[K, V]
	mut      sync.Mutex
}

func NewCacheStruct[K comparable, V any](cap int) *Cache[K, V] {
	buf := make(map[K]*Node[K, V])
	return &Cache[K, V]{buf: buf, capacity: cap}
}

func (cache *Cache[K, V]) Set(key K, val V) {
	cache.mut.Lock()         // образует потокобезопасность (блокирует доступ других горутин)
	defer cache.mut.Unlock() // в конце по любому разлочит

	if node, ok := cache.buf[key]; ok {
		MoveNode(cache, node)
		node.val = val
	} else {
		node := &Node[K, V]{prev: nil, next: nil, val: val, key: key}
		if len(cache.buf) < cache.capacity {
			AddNode(node, cache)
			cache.buf[key] = node
		} else {
			RemoveNode(cache.tail, cache)
			AddNode(node, cache)
			cache.buf[key] = node

		}
	}
}

func (cache *Cache[K, V]) Get(key K) (V, bool) {
	cache.mut.Lock()         // образует потокобезопасность (блокирует доступ других горутин)
	defer cache.mut.Unlock() // в конце по любому разлочит

	if node, ok := cache.buf[key]; ok {
		MoveNode(cache, node)
		return node.val, true
	} else {
		var zeroValue V
		return zeroValue, false
	}
}

func (cache *Cache[K, V]) Clear() {
	cache.mut.Lock()         // образует потокобезопасность (блокирует доступ других горутин)
	defer cache.mut.Unlock() // в конце по любому разлочит

	cache.head = nil
	cache.tail = nil
	clear(cache.buf)
}
