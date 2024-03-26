func romanToInt(s string) (ans int) {
    m := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    n := len(s)
    for i := range s {
        if s[i] == 'I' && i + 1 < n && (s[i+1] == 'V' || s[i+1] == 'X') {
            ans -= m[s[i]]
        } else if s[i] == 'X' && i + 1 < n && (s[i+1] == 'L' || s[i+1] == 'C') {
            ans -= m[s[i]]
        } else if s[i] == 'C' && i + 1 < n && (s[i+1] == 'D' || s[i+1] == 'M') {
            ans -= m[s[i]]
        } else {
            ans += m[s[i]]
        }
    }

    return
}