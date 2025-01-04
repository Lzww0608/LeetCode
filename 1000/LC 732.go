
import "github.com/emirpasic/gods/trees/redblacktree"
type MyCalendarThree struct {
    m *redblacktree.Tree
}


func Constructor() MyCalendarThree {
    m := redblacktree.NewWithIntComparator()
    return MyCalendarThree{m}
}


func (c *MyCalendarThree) Book(startTime int, endTime int) (ans int) {
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

    it := c.m.Iterator()
    cnt := 0
    for it.Next() {
        cnt += it.Value().(int)
        ans = max(ans, cnt)
    }
    return
}


/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */