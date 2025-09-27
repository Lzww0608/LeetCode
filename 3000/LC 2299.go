func strongPasswordCheckerII(password string) bool {
    if len(password) < 8 {
        return false
    }
    m := "!@#$%^&*()-+";
    mask := 0
    for i, c := range password {
        if c >= 'a' && c <= 'z' {
            mask |= 1
        } else if c >= 'A' && c <= 'Z' {
            mask |= 2
        } else if c >= '0' && c <= '9' {
            mask |= 4
        } else if strings.Contains(m, string(c)) {
            mask |= 8
        }

        if i > 0 && password[i] == password[i - 1] {
            return false
        }
    }

    return mask == 15
}