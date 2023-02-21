package a

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*ListNode
	head     *ListNode
	tail     *ListNode
}

type ListNode struct {
	key   int
	value int
	prev  *ListNode
	next  *ListNode
}

func New(capacity int) LRUCache {
	lru := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*ListNode),
		head:     &ListNode{key: 0, value: 0},
		tail:     &ListNode{key: 0, value: 0},
	}
	lru.head.prev = lru.tail
	lru.tail.next = lru.head
	return lru
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.cache[key]; ok {
		c.move2Head(v)
		return v.value
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if v, ok := c.cache[key]; ok {
		v.value = value
		c.move2Head(v)
	} else {
		node := &ListNode{
			key:   key,
			value: value,
		}
		c.cache[key] = node
		c.add2Head(node)
		c.size++
		if c.size > c.capacity {
			removed := c.removeTail()
			delete(c.cache, removed.key)
			c.size--
		}
	}
}

func (c *LRUCache) add2Head(node *ListNode) {
	node.prev = c.head
	node.next = c.head.next
	// DO NOT CHANGE THE ORDER
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) move2Head(node *ListNode) {
	c.removeNode(node)
	c.add2Head(node)
}

// return is the removed tail node
func (c *LRUCache) removeTail() *ListNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
