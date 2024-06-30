type MyCalendar struct {
    *redblacktree.Tree

}


func Constructor() MyCalendar {
    t := redblacktree.NewWithIntComparator()
    t.Put(math.MaxInt32, nil)
    return MyCalendar{t}
}


func (c MyCalendar) Book(start int, end int) bool {
    node, _ := c.Ceiling(end)
    it := c.IteratorAt(node)
    if !it.Prev() || it.Value().(int) <= start {
        c.Put(start, end)
        return true
    }
    return false
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
