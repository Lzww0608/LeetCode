type Node struct {
    Val int
    Next, Down *Node
}

type Skiplist struct {
    Head *Node
}


func Constructor() Skiplist {
    return Skiplist{Head: &Node{-1, nil, nil}}
}


func (this *Skiplist) Search(target int) bool {
    cur := this.Head
    for cur != nil {
        for cur.Next != nil && cur.Next.Val < target {
            cur = cur.Next
        }
        if cur.Next != nil && cur.Next.Val == target {
            return true
        }
        cur = cur.Down
    }
    return false
}


func (this *Skiplist) Add(num int)  {
    cur, isInsert := this.Head, true
    down := &Node{-1, nil, nil}
    q := []*Node{}
    for cur != nil {
        for cur.Next != nil && cur.Next.Val < num {
            cur = cur.Next
        }
        q = append(q, cur)
        cur = cur.Down
    }

    for isInsert && len(q) > 0 {
        cur = q[len(q)-1]
        q = q[:len(q)-1]
        if down.Val == -1 {
            cur.Next = &Node{num, cur.Next, nil}
        } else {
            cur.Next = &Node{num, cur.Next, down}
        }
        down = cur.Next
        isInsert = rand.Float64() < 0.5
    }

    if isInsert {
        this.Head = &Node{-1, nil, this.Head}
    }
}


func (this *Skiplist) Erase(num int) bool {
    cur, isFound := this.Head, false
    for cur != nil {
        for cur.Next != nil && cur.Next.Val < num {
            cur = cur.Next
        }
        if cur.Next != nil && cur.Next.Val == num {
            isFound = true
            cur.Next = cur.Next.Next
        }
        cur = cur.Down
    }
    return isFound
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */