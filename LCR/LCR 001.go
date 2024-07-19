func divide(a int, b int) int {
    if a == 0 {
        return 0
    } else if b == 1 {
        return a
    } else if b == -1 {
        if a == math.MinInt32 {
            return math.MaxInt32
        }
        return -a
    } else if b == 2 || b == -2 {
        return a >> 1
    }

    neg := false
    if a > 0 && b < 0 || a < 0 && b > 0 {
        neg = true
    }

    ans := plus(abs(a), abs(b))
    if neg {
        return -ans
    }

    return ans
}

func plus(a, b int) int {
    if a < b {
        return 0
    }
    
    cnt, sum := 1, b
    for sum + sum <= a {
        sum += sum
        cnt += cnt
    }

    return cnt + plus(a - sum, b)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}