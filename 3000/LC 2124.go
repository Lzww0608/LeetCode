func checkString(s string) bool {
    f := false
    for _, c := range s {
        if c == 'b' {
            f = true
        } else if c == 'a' && f {
            return false
        }
    }

    return true
}