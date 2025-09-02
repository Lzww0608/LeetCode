func longestCommonPrefix(s string, t string) int {
    i := 0

    for i < len(s) && i < len(t) {
        if s[i] != t[i] {
            i += solve(s[i + 1:], t[i:])
            break
        }
        i++
    }

    return i
}

func solve(s, t string) int {
    i := 0
    for i < len(s) && i < len(t) {
        if s[i] != t[i] {
            break
        }
        i++
    }

    return i
}