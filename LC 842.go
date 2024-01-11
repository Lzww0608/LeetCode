func splitIntoFibonacci(num string) []int {
    n, ans := len(num), []int{}
    var dfs func(a, b, idx int, res []int) bool
    dfs = func(a, b, idx int, res []int) bool {
        if idx == n {
            if len(res) >= 3 {
                ans = append(ans, res...)
                return true
            }
            return false
        }

        t := 0
        for i := idx; i < n; i++ {
            if i > idx && num[idx] == '0' {
                break
            }

            t = t*10 + int(num[i]-'0')
            // Check if t exceeds the maximum value of int
            if t > math.MaxInt32 {
                break
            }

            if len(res) < 2 || a+b == t {
                res = append(res, t)
                if dfs(b, t, i+1, res) {
                    return true
                }
                res = res[:len(res)-1]
            }
        }
        return false
    }

    dfs(-1, -1, 0, []int{})
    return ans
}
