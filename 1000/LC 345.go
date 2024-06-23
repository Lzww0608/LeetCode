func reverseVowels(t string) string {
    n := len(t)
    s := []byte(t)
    m := map[byte]bool {
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
        'A': true,
        'E': true,
        'I': true,
        'O': true,
        'U': true,
    }
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        for i < j && !m[s[i]] {
            i++
        }
        for i < j && !m[s[j]] {
            j--
        }
        s[i], s[j] = s[j], s[i]
    }
    
    return string(s)
}
