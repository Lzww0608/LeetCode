func strWithout3a3b(a int, b int) string {
    builder := strings.Builder{}

    cntA, cntB := 0, 0
    for a > 0 || b > 0 {
        if a > b && cntA < 2 || cntB == 2 {
            cntB = 0
            cntA++
            builder.WriteByte('a')
            a--
        } else {
            cntA = 0
            cntB++
            builder.WriteByte('b')
            b--
        }
    }
    
    return builder.String()
}
