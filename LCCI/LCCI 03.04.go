type MyQueue struct {
    st1 []int
    st2 []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    a, b := []int{}, []int{}
    return MyQueue{a, b}
}


/** Push element x to the back of queue. */
func (c *MyQueue) Push(x int)  {
    c.st1 = append(c.st1, x)
    return
}


/** Removes the element from in front of queue and returns that element. */
func (c *MyQueue) Pop() int {
    if len(c.st2) == 0 {
        for len(c.st1) > 0 {
            t := c.st1[len(c.st1)-1]
            c.st1 = c.st1[:len(c.st1)-1]
            c.st2 = append(c.st2, t)
        }
    }
    ans := c.st2[len(c.st2)-1]
    c.st2 = c.st2[:len(c.st2)-1]
    return ans
}


/** Get the front element. */
func (c *MyQueue) Peek() int {
    if len(c.st2) == 0 {
        for len(c.st1) > 0 {
            t := c.st1[len(c.st1)-1]
            c.st1 = c.st1[:len(c.st1)-1]
            c.st2 = append(c.st2, t)
        }
    }

    return c.st2[len(c.st2)-1]
}


/** Returns whether the queue is empty. */
func (c *MyQueue) Empty() bool {
    return len(c.st1) == 0 && len(c.st2) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */