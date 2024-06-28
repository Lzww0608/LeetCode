func shortestCompletingWord(licensePlate string, words []string) string {

    toLower := func(c byte) byte {
        if c >= 'A' && c <= 'Z' {
            return byte(c + 32)
        }
        return c
    }

    m := map[byte]int{}
    for i := range licensePlate {
        c := toLower(licensePlate[i])
        if c >= 'a' && c <= 'z' {
            m[c]++
        }
    }

    idx, d := 0, math.MaxInt32
    for i, s := range words {
        cnt := map[byte]int{}
        for i := range s {
            c := toLower(s[i])
            if c >= 'a' && c <= 'z' {
                cnt[c]++
            }
        }
        f := true
        for k, v := range m {
            if cnt[k] < v {
                f = false
                break
            }
        }
        if f && len(s) < d {
            d = len(s)
            idx = i
        }
    }

    return words[idx]
}
