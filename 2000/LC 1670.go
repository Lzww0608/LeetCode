import "container/list"

type FrontMiddleBackQueue struct {
    l, r *list.List
}

func (c *FrontMiddleBackQueue) Balance() {
    if c.l.Len() > c.r.Len() {
        val := c.l.Remove(c.l.Back())
        c.r.PushFront(val)
    }

    if c.r.Len() > c.l.Len()+1 {
        val := c.r.Remove(c.r.Front())
        c.l.PushBack(val)
    }
}

func Constructor() FrontMiddleBackQueue {
    return FrontMiddleBackQueue{l: list.New(), r: list.New()}
}

func (c *FrontMiddleBackQueue) PushFront(val int) {
    c.l.PushFront(val)
    c.Balance()
}

func (c *FrontMiddleBackQueue) PushMiddle(val int) {
    c.l.PushBack(val)
    c.Balance()
}

func (c *FrontMiddleBackQueue) PushBack(val int) {
    c.r.PushBack(val)
    c.Balance()
}

func (c *FrontMiddleBackQueue) PopFront() int {
    if c.l.Len() + c.r.Len() == 0 {
        return -1
    }

    var val interface{}
    if c.l.Len() > 0 {
        val = c.l.Remove(c.l.Front())
    } else {
        val = c.r.Remove(c.r.Front())
    }

    c.Balance()
    return val.(int)
}

func (c *FrontMiddleBackQueue) PopMiddle() int {
    if c.l.Len() + c.r.Len() == 0 {
        return -1
    }

    var val interface{}
    if c.l.Len() == c.r.Len() {
        val = c.l.Remove(c.l.Back())
    } else {
        val = c.r.Remove(c.r.Front())
    }

    c.Balance() 
    return val.(int)
}

func (c *FrontMiddleBackQueue) PopBack() int {
    if c.l.Len() + c.r.Len() == 0 {
        return -1
    }

    val := c.r.Remove(c.r.Back())
    c.Balance()
    return val.(int)
}