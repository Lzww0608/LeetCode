func countDistinct(n int64) int64 {
    ans := 0
    s := strconv.Itoa(int(n))
    m := len(s)

    pre := 1
    for i := 1; i < m; i++ {
        ans += pre * 9
        pre = pre * 9
    }

    for i := range m {
        x := int(s[i] - '0')
        if x > 0 {
            ans += pre * (x - 1)
            pre /= 9
        } else {
            break
        } 
    }    
    if !strings.Contains(s, "0") {
        ans++
    }

    return int64(ans)
}