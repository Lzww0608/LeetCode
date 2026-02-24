type Node struct {
    cnt int 
    children [2]*Node
}

func (n *Node) insert(x int) {
    cur := n 
    for i := 30; i >= 0; i-- {
        k := (x >> i) & 1 
        if cur.children[k] == nil {
            cur.children[k] = new(Node)
        }
        cur = cur.children[k]
        cur.cnt++
    }
}

func (n *Node) query(v, k int) (cnt int) {
    cur := n
    for i := 30; i >= 0 && cur != nil; i-- {
        b := (v >> i) & 1 
        if (k >> i) & 1 == 1 {
            if cur.children[b] != nil {
                cnt += cur.children[b].cnt
            }
            b ^= 1
        }

        cur = cur.children[b]
    }
    return
}

func countXorSubarrays(nums []int, k int) int64 {
    n := len(nums)
    ans := n * (n + 1) / 2

    sum := 0
    t := &Node{}
    t.insert(0)
    for _, x := range nums {
        sum ^= x
        ans -= t.query(sum, k)
        t.insert(sum)
    }

    return int64(ans)
}