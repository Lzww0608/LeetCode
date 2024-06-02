package main

import "fmt"

type Node struct {
	key       string
	val, freq int
	prev, next *Node
}

func newNode(key string) *Node {
	return &Node{
		key: key,
		val: 1,
		freq: 1,
		prev: nil,
		next: nil,
	}
}

func (c *AllOne) remove(x *Node) {
	x.next.prev = x.prev
	x.prev.next = x.next
}

func (c *AllOne) newList() *Node {
	dummy := newNode("")
	dummy.next = dummy
	dummy.prev = dummy
	return dummy
}

func (c *AllOne) pushToFront(freq int, x *Node) {
	it, found := c.freqToList[freq]
	if !found {
		it = c.newList()
		c.freqToList[freq] = it
	}
	x.next = it.next
	x.prev = it
	x.prev.next = x
	x.next.prev = x
}

func (c *AllOne) getNode(key string, val int) *Node {
	it, found := c.keyToNode[key]
	if !found {
		node := newNode(key)
		c.min_freq = 1
		c.keyToNode[key] = node
		c.pushToFront(1, node)
		if c.max_freq == 0 {
			c.max_freq = 1
		}
		return node
	}
	node := it
	c.remove(node)
	dummy := c.freqToList[node.freq]
	if dummy.next == dummy {
		delete(c.freqToList, node.freq)
		if c.min_freq == node.freq {
			c.updateMinFreq()
		}
		if c.max_freq == node.freq {
			c.max_freq--
		}
	}
	node.freq += val
	if node.freq == 0 {
		delete(c.keyToNode, key)
		return nil
	}
	c.pushToFront(node.freq, node)
	if node.freq > c.max_freq {
		c.max_freq = node.freq
	}
	if node.freq < c.min_freq || c.min_freq == 0 {
		c.min_freq = node.freq
	}
	return node
}

func (c *AllOne) updateMinFreq() {
	c.min_freq = 0
	for freq := range c.freqToList {
		if c.min_freq == 0 || freq < c.min_freq {
			c.min_freq = freq
		}
	}
}

type AllOne struct {
	min_freq, max_freq int
	keyToNode map[string]*Node
	freqToList map[int]*Node
}

func Constructor() AllOne {
	return AllOne{
		min_freq: 0,
		max_freq: 0,
		keyToNode: make(map[string]*Node),
		freqToList: make(map[int]*Node),
	}
}

func (c *AllOne) Inc(key string) {
	c.getNode(key, 1)
}

func (c *AllOne) Dec(key string) {
	c.getNode(key, -1)
	if len(c.keyToNode) == 0 {
		c.min_freq = 0
		c.max_freq = 0
	}
}

func (c *AllOne) GetMaxKey() string {
	if len(c.keyToNode) == 0 {
		return ""
	}
	it, _ := c.freqToList[c.max_freq]
	if it == nil || it.next == it {
		return ""
	}
	return it.next.key
}

func (c *AllOne) GetMinKey() string {
	if len(c.keyToNode) == 0 {
		return ""
	}
	it, _ := c.freqToList[c.min_freq]
	if it == nil || it.next == it {
		return ""
	}
	return it.next.key
}

