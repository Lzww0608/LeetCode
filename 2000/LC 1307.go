func isSolvable(words []string, result string) bool {
    i2c := make([]byte, 10)
    c2i := make([]int, 26)

    maxLen := 0
    for _, s := range words {
        if len(s) > len(result) {
            return false
        }
        maxLen = max(maxLen, len(s))
    }    
    if maxLen < len(result) - 1 {func isSolvable(words []string, result string) bool {
    i2c := make([]byte, 10)
    c2i := make([]int, 26)

    maxLen := 0
    for _, s := range words {
        if len(s) > len(result) {
            return false
        }
        maxLen = max(maxLen, len(s))
    }    
    if maxLen < len(result) - 1 {
        return false
    }

    for i := range i2c {
        i2c[i] = '#'
    }
    for i := range c2i {
        c2i[i] = -1
    }

    for i := range words {
        words[i] = reverse(words[i])
    }
    result = reverse(result)


    var solve func(int, int, int) bool
    solve = func(idx, digit, sum int) bool {
        if digit == len(result) {
            return sum == 0 && (len(result) == 1 || c2i[result[len(result)-1]-'A'] != 0)
        }
        
        if idx == len(words) {
            if c2i[result[digit] - 'A'] != -1 {
                if (sum % 10 == c2i[result[digit] - 'A']) {
                    return solve(0, digit + 1, sum / 10)
                }
            } else if i2c[sum % 10] == '#' {
                c2i[result[digit] - 'A'] = sum % 10
                i2c[sum % 10] = result[digit]
                if solve(0, digit + 1, sum / 10) {
                    return true
                }
                c2i[result[digit] - 'A'] = -1
                i2c[sum % 10] = '#'
            }
            return false
        }

        if digit >= len(words[idx]) {
            return solve(idx + 1, digit, sum)
        }

        if t := c2i[words[idx][digit] - 'A']; t != -1 {
            if t == 0 && digit == len(words[idx]) - 1 && len(words[idx])> 1 {
                return false
            }
            return solve(idx + 1, digit, sum + t)
        }

        for i := 0; i < 10; i++ {
            if i2c[i] != '#' {
                continue
            }
            if i == 0 && digit == len(words[idx]) - 1 && len(words[idx]) > 1 {
                continue
            }
            i2c[i] = words[idx][digit]
            c2i[words[idx][digit]-'A'] = i 
            if solve(idx + 1, digit, sum + i) {
                return true
            }
            i2c[i] = '#'
            c2i[words[idx][digit]-'A'] = -1
        }

        return false
    }

    return solve(0, 0, 0)
}

func reverse(s string) string {
    t := []byte(s)
    n := len(t)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        t[i], t[j] = t[j], t[i]
    }
    return string(t)
}
        return false
    }

    for i := range i2c {
        i2c[i] = '#'
    }
    for i := range c2i {
        c2i[i] = -1
    }

    for i := range words {
        words[i] = reverse(words[i])
    }
    result = reverse(result)


    var solve func(int, int, int) bool
    solve = func(idx, digit, sum int) bool {
        if digit == len(result) {
            return sum == 0
        }
        
        if idx == len(words) {
            if c2i[result[digit] - 'A'] != -1 {
                if (sum % 10 == c2i[result[digit] - 'A']) {
                    if sum % 10 == 0 && digit == len(result) - 1 && len(result) > 1 {
                        return false
                    }
                    return solve(0, digit + 1, sum / 10)
                }
            } else if i2c[sum % 10] == '#' {
                c2i[result[digit] - 'A'] = sum % 10
                i2c[sum % 10] = result[digit]
                if solve(0, digit + 1, sum / 10) {
                    return true
                }
                c2i[result[digit] - 'A'] = -1
                i2c[sum % 10] = '#'
            }
            return false
        }

        if digit >= len(words[idx]) {
            return solve(idx + 1, digit, sum)
        }

        if t := c2i[words[idx][digit] - 'A']; t != -1 {
            if t == 0 && digit == len(words[idx]) - 1 && len(words[idx])> 1 {
                return false
            }
            return solve(idx + 1, digit, sum + t)
        }

        for i := 0; i < 10; i++ {
            if i2c[i] != '#' {
                continue
            }
            if i == 0 && digit == len(words[idx]) - 1 && len(words[idx]) > 1 {
                continue
            }
            i2c[i] = words[idx][digit]
            c2i[words[idx][digit]-'A'] = i 
            if solve(idx + 1, digit, sum + i) {
                return true
            }
            i2c[i] = '#'
            c2i[words[idx][digit]-'A'] = -1
        }

        return false
    }

    return solve(0, 0, 0)
}

func reverse(s string) string {
    t := []byte(s)
    n := len(t)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        t[i], t[j] = t[j], t[i]
    }
    return string(t)
}