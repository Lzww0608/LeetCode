func printVertically(str string) []string {
    n := 0
    s := strings.Split(str, " ")
    for _, t := range s {
        n = max(n, len(t))
    }


    ans := make([]string, n)
    for i := 0; i < n; i++ {
        path := []byte{}
        for _, t := range s {
            if i >= len(t) {
                path = append(path, ' ')
            } else {
                path = append(path, t[i])
            }
        }
        tmp := string(path)
        ans[i] = strings.TrimRight(tmp, " ")
    }

    return ans
}