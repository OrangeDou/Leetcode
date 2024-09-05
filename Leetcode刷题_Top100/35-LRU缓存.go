package main

// 缓存结构体
type LRUCache struct {
	capacity   int                  //总共缓存容量
	size       int                  //当前已经存的容量
	cache      map[int]*DLinkedNode //
	head, tail *DLinkedNode         //伪头尾指针
}

// 双链表结构体
type DLinkedNode struct {
	prev  *DLinkedNode //前驱节点
	next  *DLinkedNode //后继节点
	key   int          //
	value int          //值
}

// 初始化双链表
func initDLinkNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// 初始化缓存
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	//元素存在
	node := this.cache[key]
	//被访问过，移动到队列首部
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		//修改当前值
		node := this.cache[key]
		//移动到首部
		node.value = value
		this.moveToHead(node)
	} else {
		//首次加入缓存，需要创建node存放值
		node := initDLinkNode(key, value)
		//加入缓存map
		this.cache[key] = node
		//第一次加入，所以是最近访问的，移动到头部
		this.addToHead(node)
		//当前容量++
		this.size++
		//判断是否超过缓存容量，如果超过，可以根据需要扩容或者淘汰
		if this.size > this.capacity {
			removed := this.removeTail()    //移除尾部最近未使用节点
			delete(this.cache, removed.key) //删除map中对应的键值对
			this.size--
		}
	}
}

// 将新加入的元素放到首部
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 移除节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 将元素移动到首部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 移除尾部最近未使用节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
