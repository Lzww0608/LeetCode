func addOperators(num string, target int) (ans []string) {
    n := len(num)
    str := []byte{}

    var dfs func(int, int, int)
    dfs = func(idx, sum, pre int) {
        if idx == n && sum == target {
            ans = append(ans, string(str))
            return
        } else if idx >= n {
            return
        }

        m := len(str)
        x := 0
        if idx > 0 {
            str = append(str, '+')
        }
        for i := idx; i < n; i++ {
            x = x * 10 + int(num[i] - '0')
            str = append(str, num[i])
            if idx == 0 {
                dfs(i + 1, x, x)
            } else {
                str[m] = '+'
                dfs(i + 1, sum + x, x)

                str[m] = '-'
                dfs(i + 1, sum - x, -x)

                str[m] = '*'
                dfs(i + 1, sum - pre + pre * x, pre * x)
            }
            if x == 0 {
                break
            }
        }
        str = str[:m]
    }
    dfs(0, 0, 0)

    return
}