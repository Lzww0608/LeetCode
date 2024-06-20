func compress(chars []byte) (ans int) {
    n := len(chars)
    if n == 1 {
        return 1
    }
    for l, r := 0, 1; r <= n; r++ {
        if r == n || chars[r] != chars[l] {
            if r - l == 1 {
                chars[ans] = chars[l]
                ans++
            } else {
                s := strconv.Itoa(r - l)
                chars[ans] = chars[l]
                ans++
                for t := range s {
                    chars[ans] = s[t]
                    ans++
                }
            }
            l = r
        }
        
    }

    return
}
