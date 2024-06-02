func licenseKeyFormatting(s string, k int) string {
    bytes := []byte{}
    for i := range s {
        if s[i] == '-' {
            continue
        }
        c := convert(s[i])
        bytes = append(bytes, c)
    }
    n, i := len(bytes), 0
    remainder := n % k 
    builder := strings.Builder{}
    for remainder != 0 {
        remainder--
        builder.WriteByte(bytes[i])
        i++
    }
    if builder.Len() != 0 && i < n {
        builder.WriteByte('-')
    }
    for t := 0; i < n; t++ {
        builder.WriteByte(bytes[i])
        i++
        if t % k == k - 1 && i < n {
            builder.WriteByte('-')
        }
    }

    return builder.String()
}

func convert(c byte) byte {
    if c >= 'a' && c <= 'z' {
        return byte(c - 32)
    }
    return c
}
