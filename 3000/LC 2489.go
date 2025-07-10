func fixedRatio(s string, num1 int, num2 int) int64 {
    // a / b = num1 / num2
    // (a - a') / (b - b') = num1 / num2 
    // (a - a') * num2 = (b - b') * num1
    // a * num2 - b * num1 = a' * num2 - b' * num1
    cnt := make(map[int]int)
    cnt[0] = 1
    ans := 0
    a, b := 0, 0 
    for _, c := range s {
        if c == '0' {
            a++
        } else {
            b++
        }
        ans += cnt[a * num2 - b * num1]
        cnt[a * num2 - b * num1]++
    }

    return int64(ans)
}