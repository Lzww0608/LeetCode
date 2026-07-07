func minOperations(s string, t string) (ans int) {
    n := len(s)
    if n == 1 && s[0] == '1' && s[0] != t[0] {
        return -1
    }

    // 10(2)  10(1)   11(2)   11(1)   01(3)  11(2)
    // 00     11      10      00      10     01

    for i := 0; i < n; i++ {
        if s[i] == t[i] {
            continue
        }

        if s[i] == '0' {
            ans++
        } else {
            j := i 
            for j < n && s[j] == '1' && s[j] != t[j] {
                j++
            }

            if (j - i) & 1 == 0 {
                ans += (j - i) / 2
            } else {
                ans += (j - i) / 2
                cur := 3
                if i - 1 >= 0 || j < n || j - i > 1 {
                    cur = min(cur, 2)
                } 
                ans += cur
            }

            i = j - 1
        }
    }

    return
}