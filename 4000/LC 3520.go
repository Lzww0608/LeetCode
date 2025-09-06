import "slices"

type BIT struct {
    f []int 
    n int
}

func newBIT(n int) *BIT {
    f := make([]int, n + 1)
    return &BIT{
        f: f,
        n: n,
    }
}

func (c *BIT) update(i, val int) {
    for ; i <= c.n; i += i & (-i) {
        c.f[i] += val
    }
}

func (c *BIT) query(i int) (res int) {
    for ; i > 0; i -= i & (-i) {
        res += c.f[i]
    }
    return
}

func minThreshold(nums []int, k int) int {
    m := make(map[int]int)
    for _, x := range nums {
        m[x]++
    }
    
    

    check := func(mid int) bool {
        a := make([]int, 0, len(m))
        for k := range m {
            a = append(a, k)
            a = append(a, k + mid)
        }
        sort.Ints(a)
        id := make(map[int]int)
        for i, x := range a {
            id[x] = i + 1
        }
        c := newBIT(len(a))
        cnt := 0
        for _, x := range nums {
            i := id[x]
            cnt += c.query(id[x + mid]) - c.query(i) 
            c.update(i, 1)
        }

        return cnt >= k
    }

    mx := slices.Max(nums) - slices.Min(nums) + 1
    l, r := 0, mx 
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    if r == mx {
        return -1
    }

    return r
}