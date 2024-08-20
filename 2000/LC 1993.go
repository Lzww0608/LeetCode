type LockingTree struct {
    locked   []int
    parent   []int
    children [][]int
}

func Constructor(parent []int) LockingTree {
    n := len(parent)
    locked := make([]int, n)
    children := make([][]int, n)
    for i := range locked {
        locked[i] = -1
    }
    for i, p := range parent {
        if p != -1 {
            children[p] = append(children[p], i)
        }
    }
    return LockingTree{locked, parent, children}
}

func (c *LockingTree) Lock(num int, user int) bool {
    if c.locked[num] == -1 {
        c.locked[num] = user
        return true
    }
    return false
}

func (c *LockingTree) Unlock(num int, user int) bool {
    if c.locked[num] == user {
        c.locked[num] = -1
        return true
    }
    return false
}

func (c *LockingTree) Upgrade(num int, user int) bool {

    ancestor := num
    for ancestor != -1 {
        if c.locked[ancestor] != -1 {
            return false
        }
        ancestor = c.parent[ancestor]
    }


    hasLockedDescendant := false
    var unlockDescendants func(int)
    unlockDescendants = func(node int) {
        for _, child := range c.children[node] {
            if c.locked[child] != -1 {
                hasLockedDescendant = true
                c.locked[child] = -1
            }
            unlockDescendants(child)
        }
    }
    unlockDescendants(num)

    if !hasLockedDescendant {
        return false
    }

    c.locked[num] = user
    return true
}
