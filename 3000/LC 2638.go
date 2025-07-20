func countTheNumOfKFreeSubsets(nums []int, k int) int64 {
    ans := 1
    a := make([][]int, k)
    for _, x := range nums {
        y := x % k 
        a[y] = append(a[y], x)
    }

    for _, v := range a {
        sort.Ints(v)
        n := len(v)
        f := make([]int, n + 1)
        f[0] = 1
        if n > 0 {
            f[1] = 2
        }
        for i := 2; i <= n; i++ {
            if v[i-1] - v[i-2] == k {
                f[i] = f[i-2] + f[i-1]
            } else {
                f[i] = f[i-1] * 2
            }
        }

        ans *= f[n]
    }

    return int64(ans)
}

