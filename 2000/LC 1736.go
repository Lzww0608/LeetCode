func maximumTime(time string) (ans string) {
    s := strings.Split(time, ":")

    if s[0][0] == '?' {
        if s[0][1] == '?' {
            ans += "23"
        } else if s[0][1] < '4' {
            ans += "2"
            ans += s[0][1:]
        } else {
            ans += "1"
            ans += s[0][1:]
        }
    } else if s[0][1] == '?' {
        if s[0][0] == '1' {
            ans += "19"
        } else if s[0][0] == '0' {
            ans += "09"
        }else {
            ans += "23"
        }
    } else {
        ans += s[0]
    }
    ans += ":"
    if s[1][0] == '?' {
        if s[1][1] == '?' {
            ans += "59"
        } else {
            ans += "5"
            ans += s[1][1:]
        }
    } else if s[1][1] == '?' {
        ans += s[1][:1]
        ans += "9"
    } else {
        ans += s[1]
    }

    return 
}