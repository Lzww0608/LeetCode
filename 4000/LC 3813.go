var t string = "aeiou"
func vowelConsonantScore(s string) int {
    v, c := 0, 0
    for i := range s {
        if strings.Contains(t, string(s[i])) {
            v++
        } else if unicode.IsLetter(rune(s[i])) {
            c++
        }
    }

    if c > 0 {
        return v / c
    }
    return 0
}