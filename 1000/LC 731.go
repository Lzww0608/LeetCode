
import "github.com/emirpasic/gods/trees/redblacktree"
type MyCalendarTwo struct {
    m *redblacktree.Tree
}


func Constructor() MyCalendarTwo {
    return MyCalendarTwo{m: redblacktree.NewWithIntComparator()}
}


func (c *MyCalendarTwo) Book(startTime int, endTime int) bool {
    if v, ok := c.m.Get(startTime); ok {
        c.m.Put(startTime, v.(int) + 1)
    } else {
        c.m.Put(startTime, 1)
    }

    if v, ok := c.m.Get(endTime); ok {
        c.m.Put(endTime, v.(int) - 1)
    } else {
        c.m.Put(endTime, -1)
    }

    cnt := 0
    it := c.m.Iterator()

    for it.Next() {
        value := it.Value()
        cnt += value.(int)
        if cnt > 2 {
            if v, ok := c.m.Get(startTime); ok {
                c.m.Put(startTime, v.(int) - 1)
            } 

            if v, ok := c.m.Get(endTime); ok {
                c.m.Put(endTime, v.(int) + 1)
            } 
            return false
        }
    }

    return true
}


/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */