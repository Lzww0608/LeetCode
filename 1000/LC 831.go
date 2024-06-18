func maskPII(s string) string {
    if strings.Contains(s, "@") {
        return email(s)
    }

    return phone(s)
}

func email(s string) string {
    str := strings.Split(s, "@")
    a, b := str[0], str[1]
    builder := strings.Builder{}

    if a[0] >= 'A' && a[0] <= 'Z' {
        builder.WriteByte(byte(a[0] + 32))
    } else {
        builder.WriteByte(a[0])
    }
    builder.WriteString("*****")
    if a[len(a)-1] >= 'A' && a[len(a)-1] <= 'Z' {
        builder.WriteByte(byte(a[len(a)-1] + 32))
    } else {
        builder.WriteByte(a[len(a)-1])
    }

    builder.WriteByte('@')
    for i := range b {
        if b[i] >= 'A' && b[i] <= 'Z' {
            builder.WriteByte(byte(b[i] + 32))
        } else {
            builder.WriteByte(b[i])
        }
    }

    return builder.String()
}

func phone(s string) string {
    a := []byte{}
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            a = append(a, s[i])
        }
    }
    if len(a) == 10 {
        return "***-***-" + string(a[6:])
    } else if len(a) == 11 {
        return "+*-***-***-" + string(a[7:])
    } else if len(a) == 12 {
        return "+**-***-***-" + string(a[8:])
    } 
    return "+***-***-***-" + string(a[9:])
}