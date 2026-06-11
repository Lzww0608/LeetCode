func generateValidStrings(n int, k int) (ans []string) {
    for x := range (1 << n) {
        s := fmt.Sprintf("%0*b", n, x)
        cost := 0
        for i := range s {
            if s[i] == '1' {
                cost += i
                if i > 0 && s[i - 1] == '1' {
                    cost = k + 1
                    break
                }
            }
        }

        if cost <= k {
            ans = append(ans, s)
        }
    }

    return
}