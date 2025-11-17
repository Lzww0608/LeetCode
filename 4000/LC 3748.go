type pair struct {
    l, r int
}
func countStableSubarrays(nums []int, queries [][]int) []int64 {
    n := len(nums)
    a := []pair{}
    for i := 0; i < n; {
        j := i + 1
        for j < n && nums[j] >= nums[j - 1] {
            j++
        }
        a = append(a, pair{i, j - 1})
        i = j
    }

    m := len(a)
    b := make([]int, m + 1)
    for i := range m {
        x := a[i].r - a[i].l + 1
        b[i + 1] = b[i] + x * (x + 1) / 2
    }
    a = append(a, pair{n, n})
    //fmt.Println(b)

    solve := func(v pair) int {
        l := sort.Search(m, func(i int) bool {
            return a[i].l >= v.l
        })
        r := sort.Search(m, func(i int) bool {
            return a[i].l > v.r
        })

        if l == r {
            x := v.r - v.l + 1
            return x * (x + 1) / 2
        }

        res := 0
        if a[l].l > v.l {
            x := a[l - 1].r - v.l + 1
            res += x * (x + 1) / 2 
        }

        res += b[r - 1] - b[l]

        x := v.r - a[r - 1].l + 1
        res += x * (x + 1) / 2

        return res
    }

    ans := make([]int64, len(queries))
    for i, q := range queries {
        if q[0] == q[1] {
            ans[i] = 1
        } else {
            ans[i] = int64(solve(pair{q[0], q[1]}))
        }
        
    }    

    return ans
}