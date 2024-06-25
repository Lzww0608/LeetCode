func findDuplicate(paths []string) (ans [][]string) {
    m := map[string]string{}
    id := map[string]int{}

    for _, path := range paths {
        s := strings.Split(path, " ")
        head := s[0]
        for i := 1; i < len(s); i++ {
            ss := strings.Split(s[i], "(")
            t := head + "/" + ss[0]
            content := ss[1][:len(ss[1])-1]
            if _, ok := m[content]; ok {
                if _, ex := id[content]; ex {
                    ans[id[content]] = append(ans[id[content]], t)
                } else {
                    ans = append(ans, []string{t, m[content]})
                    n := len(ans) - 1
                    id[content] = n
                }
            } else {
                m[content] = t
            }
        }
    }

    return
}
