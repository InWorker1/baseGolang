package main

type Node[K comparable, V any] struct {
	prev *Node[K, V]
	next *Node[K, V]
	val  V
	key  K
}

func MoveNode[K comparable, V any](cache *Cache[K, V], node *Node[K, V]) {
	if node == nil || cache == nil || node.prev == nil {
		return
	} else if cache.head == nil {
		cache.head = node
		cache.tail = node
		return
	}
	if node.next != nil {
		node.next.prev = node.prev
		node.prev.next = node.next
	} else {
		node.prev.next = nil
		cache.tail = node.prev
	}
	node.next = cache.head
	node.prev = nil
	cache.head.prev = node
	cache.head = node
}

func RemoveNode[K comparable, V any](node *Node[K, V], cache *Cache[K, V]) {
	if node == nil || cache == nil {
		return
	}
	if node == cache.head && node == cache.tail {
		cache.tail = nil
		cache.head = nil
	} else if node == cache.head {
		cache.head = node.next
		cache.head.prev = nil
	} else if node == cache.tail {
		cache.tail = node.prev
		cache.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.next = nil
	node.prev = nil
	delete(cache.buf, node.key)
}

func AddNode[K comparable, V any](node *Node[K, V], cache *Cache[K, V]) {
	if node == nil || cache == nil {
		return
	}
	if cache.head == nil && cache.tail == nil {
		node.next = nil
		node.prev = nil
		cache.head = node
		cache.tail = node

	} else {
		cache.head.prev = node
		node.next = cache.head
		cache.head = node
	}
}
