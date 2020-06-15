package leetcode

// https://leetcode-cn.com/problems/lru-cache/
// 146. LRU缓存机制

type LRUCache struct {
	capacity int
	cacheMap map[int]*Node
	first    *Node
	last     *Node
}

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

func newNode(k, v int) *Node {
	return &Node{
		key: k,
		val: v,
	}

}

func Constructor(capacity int) LRUCache {
	first := newNode(0, 0)
	last := newNode(0, 0)
	first.next = last
	last.prev = first
	return LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]*Node),
		first:    first,
		last:     last,
	}
}

func (this *LRUCache) Get(key int) int {
	var res int
	if v, ok := this.cacheMap[key]; !ok {
		return -1
	} else {
		this.remove(v)
		this.InsertFirst(v)
		res = v.val
	}
	return res
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cacheMap[key]; ok {
		v.val = value
		if v != this.first.next {
			this.remove(v)
			this.InsertFirst(v)
		}
	} else {
		if len(this.cacheMap) == this.capacity {
			delete(this.cacheMap, this.last.prev.key)
			this.remove(this.last.prev)
		}
		node := newNode(key, value)
		this.cacheMap[key] = node
		this.InsertFirst(node)
	}
}

func (this *LRUCache) InsertFirst(node *Node) {
	node.next = this.first.next
	this.first.next = node

	node.prev = this.first
	node.next.prev = node
}

func (this *LRUCache) remove(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}
