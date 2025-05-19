package main

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LRUCache struct {
	Capacity  int
	Map       map[int]*Node
	DummyHead *Node
	DummyTail *Node
}

func Constructor(capacity int) *LRUCache {
	cache := &LRUCache{
		Capacity:  capacity,
		Map:       make(map[int]*Node),
		DummyHead: &Node{},
		DummyTail: &Node{},
	}
	cache.DummyHead.Next = cache.DummyTail
	cache.DummyTail.Pre = cache.DummyHead
	return cache
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.Map[key]; ok {
		return -1
	} else {
		// remove
		c.removeNode(node)
		// update
		c.updateHead(node)
		return node.Value
	}
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.Map[key]; ok {
		node.Value = value
		// remove

		// uodate

	} else {
		node = &Node{Key: key, Value: value}
		c.Map[key] = node
		// update
		c.updateHead(node)
		if len(c.Map) > c.Capacity {
			lastNode := c.DummyTail.Pre
			// remove
			c.removeNode(lastNode)
			delete(c.Map, lastNode.Key)
		}
	}
}

func (c *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (c *LRUCache) updateHead(node *Node) {
	node.Next = c.DummyHead.Next
	node.Pre = c.DummyHead
	c.DummyHead.Next.Pre = node
	c.DummyHead.Next = node
}
