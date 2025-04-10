
import "strconv"
func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
    high := strconv.Itoa(int(finish))
    low := strconv.Itoa(int(start))
    low = strings.Repeat("0", len(high) - len(low)) + low
    n := len(high)
    dif := n - len(s)
    memo := make([]int, n)
    for i := range memo {
        memo[i] = -1
    }

    var dfs func(int, bool, bool) int 
    dfs = func(i int, isLow, isHigh bool) (ans int) {
        if i == n {
            return 1
        }

        if !isLow && !isHigh && memo[i] != -1 {
            return memo[i]
        }

        l, r := 0, 9
        if isLow {
            l = int(low[i] - '0') 
        }
        if isHigh {
            r = int(high[i] - '0')
        }

        if i < dif {
            for d := l; d <= min(r, limit); d++ {
                ans += dfs(i + 1, isLow && d == l, isHigh && d == r)
            }
        } else {
            x := int(s[i-dif] - '0')
            if x >= l && x <= r {
                ans += dfs(i + 1, isLow && x == l, isHigh && x == r)
            }
        }

        if !isLow && !isHigh {
            memo[i] = ans
        }

        return
    }

    return int64(dfs(0, true, true))
}