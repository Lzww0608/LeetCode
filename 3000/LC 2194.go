func cellsInRange(s string) (ans []string) {
    a := strings.Split(s, ":")
    for c := a[0][0]; c <= a[1][0]; c++ {
        for r := a[0][1]; r <= a[1][1]; r++ {
            ans = append(ans, fmt.Sprintf("%c%c", c, r))
        }
    }

    return
}