func isAdditiveNumber(num string) bool {
    n := len(num)
    if n < 3 {
        return false
    }

    var dfs func(pre, sum, i int) bool
    dfs = func(pre, sum, i int) bool {
        if i == n {
            return true
        }
        cur := 0
        for j := i; j < n; j++ {
            if num[i] == '0' && j > i {
                break
            }
            cur = cur * 10 + int(num[j] - '0')
            if cur > sum {
                break
            } else if cur == sum && dfs(sum, pre + sum, j + 1) {
                return true
            }
        }

        return false
    }

    for i := 0; i < n - 2; i++ {
        if num[0] == '0' && i > 0 {
            break
        } 
        a, _ := strconv.Atoi(num[:i+1])
        for j := i + 1; j < n - 1; j++ {
            if num[i+1] == '0' && j > i + 1 {
                break
            }
            b, _ := strconv.Atoi(num[i+1:j+1])
            if dfs(b, a + b, j + 1) {
                return true
            }
        }
    }

    return false
}