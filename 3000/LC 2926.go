
import "math"
type BIT []int 

func (b BIT) update(i, val int) {
    for ; i < len(b); i += i & -i {
        b[i] = max(b[i], val)
    }
}

func (b BIT) query(i int) int {
    ans := math.MinInt
    for ; i > 0; i &= i - 1 {
        ans = max(ans, b[i])
    }

    return ans
}

func maxBalancedSubsequenceSum(nums []int) int64 {
    b := slices.Clone(nums)
    for i := range b {
        b[i] -= i 
    }

    slices.Sort(b)
    b = slices.Compact(b)

    t := make(BIT, len(b) + 1)
    for i := range t {
        t[i] = math.MinInt
    }

    for i, x := range nums {
        j := sort.SearchInts(b, x - i) + 1 
        f := max(t.query(j), 0) + x 
        t.update(j, f)
    }

    return int64(t.query(len(b)))
}