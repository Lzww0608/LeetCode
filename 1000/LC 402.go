func removeKdigits(a string, k int) string {
	s := []byte{}

    for i := range a {
        c := a[i]
        for k > 0 && len(s) > 0 && s[len(s)-1] > c {
            k--
            s = s[:len(s)-1]
        }
        if c != '0' || len(s) != 0 {
            s = append(s, c)
        }
    }

    for k > 0 && len(s) > 0 {
        k--
        s = s[:len(s)-1]
    }

    if len(s) == 0 {
        return "0"
    }

    return string(s)
}
