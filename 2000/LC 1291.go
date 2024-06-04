func sequentialDigits(low int, high int) []int {
    s1, s2 := strconv.Itoa(low), strconv.Itoa(high)
    n, m := len(s1), len(s2)
    ans := []int{}
    for i := n; i <= m; i++ {
        t := make([]rune, i)
        for j := 1; j + i - 1 < 10; j++ {
            for k := 0; k < i; k++ {
                t[k] = rune(k + j + '0')
            }
            res, _ := strconv.Atoi(string(t))
            if res >= low && res <= high {
                ans = append(ans, res)
            } 
        }
    }
    return ans
}
