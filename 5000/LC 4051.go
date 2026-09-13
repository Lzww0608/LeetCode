func distantSubarrays(nums []int, goal int, k int) int64 {
    m := make(map[int]bool)
    sum := 0
    m[0] = true

    for _, x := range nums {
        sum += x
        m[sum] = true
    }

    a := make([]int, 0, len(m))
    for x := range m {
        a = append(a, x)
    }

    sort.Ints(a)
    pos := make(map[int]int)
    for i, x := range a {
        pos[x] = i
    }

    n := len(a)
    f := NewBIT(n)
    f.update(pos[0] + 1)

    sum = 0
    var ans int64
    for s, x := range nums {
        sum += x

        low := sum - goal - k
        high := sum - goal + k
        l := sort.Search(len(a), func(i int) bool {
            return a[i] > low
        })

        left := f.query(l)
        r := sort.SearchInts(a, high)
        right := s + 1 - f.query(r)

        if k == 0 {
            ans += int64(s + 1)
        } else {
            ans += int64(left + right)
        }

        f.update(pos[sum] + 1)
    }

    return ans
}

type BIT struct {
    n int
    f []int
}

func NewBIT(n int) BIT {
    return BIT{
        n: n,
        f: make([]int, n+1),
    }
}

func (b BIT) update(i int) {
    for i <= b.n {
        b.f[i]++
        i += i & -i
    }
}

func (b BIT) query(i int) (res int) {
    for i > 0 {
        res += b.f[i]
        i -= i & -i
    }
    return
}