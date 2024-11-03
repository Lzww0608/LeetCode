
func numberOfRounds(loginTime string, logoutTime string) int {
    var h1, h2, m1, m2 int
    fmt.Sscanf(loginTime, "%d:%d", &h1, &m1)
    fmt.Sscanf(logoutTime, "%d:%d", &h2, &m2)
    if loginTime > logoutTime {
        h2 += 24
    }

    s, t := h1 * 60 + m1, h2 * 60 + m2
    return (t - t % 15 - s) / 15
}