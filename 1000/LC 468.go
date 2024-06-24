func validIPAddress(queryIP string) string {
    if judgeIPv4(queryIP) {
        return "IPv4"
    } else if judgeIPv6(queryIP) {
        return "IPv6"
    }

    return "Neither"
}

func judgeIPv4(IP string) bool {
    str := strings.Split(IP, ".")
    if len(str) != 4 {
        return false
    }
    for _, s := range str {
        if len(s) <= 0 || len(s) > 3 || s[0] == '0' && len(s) > 1 {
            return false
        }
        x := 0
        for _, c := range s {
            if c >= '0' && c <= '9' {
                x = x * 10 + int(c - '0')
            } else {
                return false
            }
        }
        if x > 255 {
            return false
        }
    }
    return true
}

func judgeIPv6(IP string) bool {
    str := strings.Split(IP, ":")
    if len(str) != 8 {
        return false
    }
    for _, s := range str {
        if len(s) > 4 || len(s) == 0 {
            return false
        }
        for _, c := range s {
            if c >= '0' && c <= '9' || c >= 'a'&& c <= 'f' || c >= 'A' && c <= 'F' {
                continue
            } else {
                return false
            }
        }
    }
    return true
}