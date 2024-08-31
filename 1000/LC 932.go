func beautifulArray(n int) []int {
    ans := make([]int, 0, n)

    var dfs func([]int)
    dfs = func(a []int) {
        n := len(a)
        if n <= 2 {
            for _, x := range a {
                ans = append(ans, x)
            }
            return
        } else if n == 3 {
            ans = append(ans, []int{a[1], a[0], a[2]}...)
            return
        }

        b := []int{}
        c := []int{}
        for i, x := range a {
            if i & 1 == 0 {
                b = append(b, x)
            } else {
                c = append(c, x)
            }
        }
        dfs(b)
        dfs(c)
    }

    a := []int{}
    b := []int{}
    for i := 1; i <= n; i++ {
        if i & 1 == 1 {
            a = append(a, i)
        } else {
            b = append(b, i)
        }
    }

    dfs(a)
    dfs(b)

    return ans
}