
type pair struct {
    score int
    name string
}

func compare(x, y any) int {
    a, b := x.(pair), y.(pair)
    if a.score > b.score || a.score == b.score && a.name < b.name {
        return -1
    }
    return 1
}

type SORTracker struct {
    *redblacktree.Tree
    cur redblacktree.Iterator
}


func Constructor() SORTracker {
    root := redblacktree.NewWith(compare)
    root.Put(pair{}, nil) 
    return SORTracker{root, root.IteratorAt(root.Left())}
}


func (c *SORTracker) Add(name string, score int)  {
    p := pair{score, name}
    c.Put(p, nil) 
    if compare(p, c.cur.Key()) < 0 {
        c.cur.Prev()
    }
}


func (c *SORTracker) Get() string {
    name := c.cur.Key().(pair).name
    c.cur.Next()
    return name
}


/**
 * Your SORTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(name,score);
 * param_2 := obj.Get();
 */