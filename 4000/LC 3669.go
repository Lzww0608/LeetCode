const N int = 100_001

var div [N][]int

func init() {
    for i := 1; i < N; i++ {
        for j := i; j < N; j += i {
            div[j] = append(div[j], i)
        }
    }
}

func minDifference(n int, k int) (ans []int) {
    a := make([]int, k)
    d := math.MaxInt32
    var dfs func(int, int, int, int)
    dfs = func(n, j, mn, mx int) {
        if j == k {
            if mx - mn > d || n != 1 {
                return
            }
            d = mx - mn
            ans = slices.Clone(a)
            return
        }

        for _, x := range div[n] {
            a[j] = x
            dfs(n / x, j + 1, min(mn, x), max(mx, x))
        }
    }

    dfs(n, 0, math.MaxInt32, 0)
    return
}