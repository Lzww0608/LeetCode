func myAtoi(s string) int {
    var ans int64 = 0
    isNegative := false
    i, n := 0, len(s)

    for i < n && s[i] == ' ' {
        i++
    }

    if i < n && s[i] == '-' {
        isNegative = true
        i++
    } else if i < n && s[i] == '+' {
        i++
    }

    for i < n && s[i] >= '0' && s[i] <= '9' {
        ans = ans * 10 + int64(s[i] - '0')
        if ans > math.MaxInt32 + 1 || (isNegative == false && ans > math.MaxInt32) {
            break
        }
        i++
    }

    if isNegative {
        ans *= -1
        ans = max(ans, math.MinInt32)
    } else {
        ans = min(ans, math.MaxInt32)
    }

    return int(ans)
}