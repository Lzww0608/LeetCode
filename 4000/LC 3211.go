
func validStrings(n int) []string {
    s := []string{"0", "1"}
    for n > 1 {
        n--
        tmp := []string{}
        for _, t := range s {
            tmp = append(tmp, t + "1")
            if t[len(t)-1] != '0' {
                tmp = append(tmp, t + "0")
            }
        }
        s = tmp
    }

    return s
}