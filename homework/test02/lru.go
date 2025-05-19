package main

// Node 节点结构体字段
type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

// LRUCache Cache结构
type LRUCache struct {
	Capacity  int
	Map       map[int]*Node
	DummyHead *Node
	DummyTail *Node
}

// Constructor 构造器实现
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity:  capacity,
		Map:       make(map[int]*Node),
		DummyHead: &Node{},
		DummyTail: &Node{},
	}
	cache.DummyHead.Next = cache.DummyTail
	cache.DummyTail.Pre = cache.DummyHead
	return cache
}

// Get 获取元素
func (c *LRUCache) Get(key int) int {
	if node, ok := c.Map[key]; ok {
		return -1
	} else {
		c.removeNode(node)
		c.updateHead(node)
		return node.Value
	}
}

// Put 更新元素
func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.Map[key]; ok {
		node.Value = value
		c.removeNode(node)
		c.updateHead(node)
	} else {
		node = &Node{Key: key, Value: value}
		c.Map[key] = node
		c.updateHead(node)
		if len(c.Map) > c.Capacity {
			lastNode := c.DummyTail.Pre
			c.removeNode(lastNode)
			delete(c.Map, lastNode.Key)
		}
	}
}

// 移除节点
func (c *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

// 更新
func (c *LRUCache) updateHead(node *Node) {
	node.Next = c.DummyHead.Next
	node.Pre = c.DummyHead
	c.DummyHead.Next.Pre = node
	c.DummyHead.Next = node
}
