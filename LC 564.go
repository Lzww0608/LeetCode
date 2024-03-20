func nearestPalindromic(s string) string {
    n := len(s)

    getPal := func(x int64) int64 {
        str := strconv.FormatInt(x, 10)
        m := len(str)
        if m < n {
            return int64(math.Pow(10, float64(m))) - 1
        }

        bytes := []byte(str)
        for i, j := 0, m - 1; i < j; i, j = i + 1, j - 1 {
            bytes[j] = bytes[i]
        }

        res, _ := strconv.ParseInt(string(bytes), 10, 64)
        return res
    }

    num, _ := strconv.ParseInt(s, 10, 64)
    ans1 := getPal(num - 1)
    ans2 := getPal(num + 1)

    if ans1 >= num {
        ans1 = getPal(num - int64(math.Pow(10, float64(n / 2))))
    }
    if ans2 <= num {
        ans2 = getPal(num + int64(math.Pow(10, float64(n / 2))))
    }
    fmt.Println(ans1, ans2)
    if num - ans1 <= ans2 - num {
        return strconv.FormatInt(ans1, 10)
    }

    return strconv.FormatInt(ans2, 10)
}