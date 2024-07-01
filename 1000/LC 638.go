func shoppingOffers(price []int, special [][]int, needs []int) int {
    n := len(price)
    dp := map[string]int{}
    var dfs func([]int) int
    dfs = func(_needs []int) int {
        key := sliceToString(_needs)
        if v, ok := dp[key]; ok {
            return v
        }
        minn := 0
        for i, x := range _needs {
            minn += price[i] * x
        }

        for _, v := range special {
            f := true
            nextNeeds := make([]int, n)
            copy(nextNeeds, _needs)
            for j := 0; j < n; j++ {
                if v[j] > nextNeeds[j] {
                    f = false
                }
                nextNeeds[j] -= v[j]
            }
            if !f {
                continue
            }
            minn = min(minn, dfs(nextNeeds) + v[n])
        }
        dp[key] = minn
        return minn
    }
    return dfs(needs)
}

func sliceToString(slice []int) string {
    str := ""
    for _, v := range slice {
        str += fmt.Sprintf("%d, ", v)
    }
    return str
}
