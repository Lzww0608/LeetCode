type Allocator struct {
    a []int
}


func Constructor(n int) Allocator {
    a := make([]int, n)
    return Allocator{a}
}


func (c *Allocator) Allocate(size int, mID int) int {
    n := len(c.a)
    cnt := 0
    for i := 0; i < n; i++ {
        if c.a[i] == 0 {
            if cnt++; cnt == size {
                for j := i - cnt + 1; j <= i; j++ {
                    c.a[j] = mID
                }
                return i - cnt + 1
            }
        } else {
            cnt = 0
        }
    }
    return -1
}


func (c *Allocator) FreeMemory(mID int) (ans int) {
    for i := range c.a {
        if c.a[i] == mID {
            c.a[i] = 0
            ans++
        }
    }

    return
}


/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */