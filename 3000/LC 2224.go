func convertTime(current string, correct string) (ans int) {
    s := strings.Split(current, ":")
    t := strings.Split(correct, ":")
    h1, _ := strconv.Atoi(s[0])
    h2, _ := strconv.Atoi(t[0])
    ans += h2 - h1

    m1, _ := strconv.Atoi(s[1])
    m2, _ := strconv.Atoi(t[1])
    if m2 >= m1 {
        dif := m2 - m1
        ans += dif / 15
        dif %= 15
        ans += dif / 5
        dif %= 5
        ans += dif / 1 
    } else {
        ans--
        dif := 60 - m1 + m2 
        ans += dif / 15
        dif %= 15
        ans += dif / 5
        dif %= 5
        ans += dif / 1 
    }

    return
}