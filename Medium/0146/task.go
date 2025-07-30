package lm0146 // https://leetcode.com/problems/lru-cache/description/

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	head     *Node         // указатель на начало списка (самый новый элемент)
	tail     *Node         // указатель на конец списка (самый старый элемент)
	data     map[int]*Node // быстрое получение значений по ключу
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*Node, capacity),
	}
}

func (lru *LRUCache) delete(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lru.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lru.tail = node.prev
	}
}

func (lru *LRUCache) add(node *Node) {
	node.next = lru.head
	node.prev = nil

	if lru.head != nil {
		lru.head.prev = node
	}
	lru.head = node

	if lru.tail == nil {
		lru.tail = node
	}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.data[key]
	if !ok {
		return -1
	}

	// если элемент всего один
	if len(lru.data) == 1 {
		return node.value
	}

	lru.delete(node)
	lru.add(node)

	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	// ключ есть
	if lru.Get(key) != -1 {
		// Обновляем значение
		lru.data[key].value = value

		return
	}

	// создаём новый элемент
	node := &Node{key: key, value: value}

	// если кол-во элементов мапы равно капасити, удаляем самый старый
	if len(lru.data) >= lru.capacity {
		// удаляем из мапы самый старый ключ (tail)
		delete(lru.data, lru.tail.key)

		lru.delete(lru.tail)
	}

	// добавляем новый элемент в мапу
	lru.data[key] = node

	// добавляем новый элемент в список
	lru.add(node)
}
