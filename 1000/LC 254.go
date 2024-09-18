func getFactors(n int) (ans [][]int) {
    var dfs func([]int)
    dfs = func(a []int) {
        ans = append(ans, a)

        m := len(a)
        for i := a[m-2]; i * i <= a[m-1]; i++ {
            if a[m-1] % i == 0 {
                tmp := append([]int(nil), a[:m-1]...)
                tmp = append(tmp, i, a[m-1] / i)
                dfs(tmp)
            }
        }
    }

    for i := 2; i * i <= n; i++ {
        if n % i == 0 {
            dfs([]int{i, n / i})
        }
    }

    return 
}