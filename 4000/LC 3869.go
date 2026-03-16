const (
    INIT = iota
    INC
    DEC
    NOT_GOOD
)

func countFancy(ll int64, rr int64) int64 {
    left := strconv.Itoa(int(ll))
    right := strconv.Itoa(int(rr))
    dif := len(right) - len(left)
    n := len(right)

    memo := make([][][4][10]int, n * 9 + 1)
    for i := range memo {
        memo[i] = make([][4][10]int, n + 1)
    }
    check := func(x int) bool {
        if x < 100 {
            return x / 10 != x % 10
        }
        a, b, c := x / 100, x / 10 % 10, x % 10

        return a < b && b < c || a > b && b > c
    }

    var dfs func(int, int, int, int, bool, bool) int 
    dfs = func(i, sum, pre, state int, isLow, isHigh bool) int {
        if i == n {
            if state != NOT_GOOD || check(sum) {
                return 1
            }
            return 0
        }

        res := memo[sum][i][state][pre]
        if !isLow && !isHigh && res != 0 {
            return res - 1
        }

        res = 0
        l, r := 0, 9 
        if isLow && i >= dif {
            l = int(left[i - dif] - '0')
        }
        if isHigh {
            r = int(right[i] - '0')
        }

        d := l  
        if isLow && i < dif {
            res = dfs(i + 1, 0, 0, INIT, true, false)
            d = 1
        }

        for ; d <= r; d++ {
            nxt := state 
            switch state {
                case INIT:
                    if pre > 0 {
                        if d > pre {
                            nxt = INC 
                        } else if d < pre {
                            nxt = DEC
                        } else {
                            nxt = NOT_GOOD
                        }
                    }  
                case INC:
                    if d <= pre {
                        nxt = NOT_GOOD
                    }
                case DEC:
                    if d >= pre {
                        nxt = NOT_GOOD
                    }
            }
            res += dfs(i + 1, sum + d, d, nxt, isLow && d == l, isHigh && d == r)
        }

        if !isHigh && !isLow {
            memo[sum][i][state][pre] = res + 1
        }
        return res
    }

    return int64(dfs(0, 0, 0, INIT, true, true))
}