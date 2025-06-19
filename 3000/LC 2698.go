func punishmentNumber(n int) (ans int) {
    for i := 1; i <= n; i++ {
        if check(i) {
            ans += i * i
        }
    }

    return 
}

func check(x int) bool {
    y := x * x 
    
    var dfs func(int, int) bool 
    dfs = func(x, need int) bool {
        if need < 0 {
            return false
        } else if x == need {
            return true
        }

        for i := 10; i < x; i *= 10 {
            t := x % i
            if dfs(x /i, need - t) {
                return true
            }
        }

        return false
    }
    

    return dfs(y, x)
}