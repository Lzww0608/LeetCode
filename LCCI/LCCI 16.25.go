type Node struct {
    key, val int
    Prev, Next *Node
}

type LRUCache struct {
    capacity int
    keyToNode map[int]*Node
    dummy *Node
}


func Constructor(capacity int) LRUCache {
    dummy := &Node{}
    dummy.Next = dummy
    dummy.Prev = dummy
    keyToNode := make(map[int]*Node)

    return LRUCache{capacity, keyToNode, dummy}
}

func remove(x *Node) {
    x.Prev.Next = x.Next
    x.Next.Prev = x.Prev
}

func (c *LRUCache) pushToFront(x *Node) {
    x.Next = c.dummy.Next
    x.Prev = c.dummy
    c.dummy.Next.Prev = x
    c.dummy.Next = x
}

func (c *LRUCache) getNode(key int) *Node {
    if node, ok := c.keyToNode[key]; ok {
        remove(node)
        c.pushToFront(node)
        return node
    }
    return nil
}

func (c *LRUCache) Get(key int) int {
    if node := c.getNode(key); node != nil {
        return node.val
    }
    return -1
}

func (c *LRUCache) Put(key int, value int) {
    if node := c.getNode(key); node != nil {
        node.val = value
        return
    }

    node := &Node{
        key: key,
        val: value,
    }
    c.pushToFront(node)
    c.keyToNode[key] = node

    if len(c.keyToNode) > c.capacity {
        last := c.dummy.Prev
        remove(last)
        delete(c.keyToNode, last.key)
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
