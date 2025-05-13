package main

import "container/list"

func main() {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// LRU 缓存属于超高频面试题，必背了属于是

type entry struct {
	key, val int
}

type LRUCache struct {
	capacity  int
	list      *list.List
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, list: list.New(), keyToNode: make(map[int]*list.Element)}
}

func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil {
		return -1
	}
	// 访问过，更新到最前面
	c.list.MoveToFront(node)
	return node.Value.(entry).val
}

func (c *LRUCache) Put(key int, value int) {
	if node := c.keyToNode[key]; node != nil {
		node.Value = entry{key, value}
		c.list.MoveToFront(node)
		return
	}
	c.keyToNode[key] = c.list.PushFront(entry{key, value})
	if len(c.keyToNode) > c.capacity {
		// 删除最后一个节点
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key)
	}
}
