func beautifulSubstrings(s string, k int) (ans int) {
    m := map[byte]bool {
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }

    n := len(s)
    for i := 0; i < n; i++ {
        a, b := 0, 0
        for j := i; j < n; j++ {
            if m[s[j]] {
                a++
            } else {
                b++
            }
            if a == b && a * b % k == 0 {
                ans++
            }
        }
    }

    return
}