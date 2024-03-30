func alertNames(keyName []string, keyTime []string) (ans []string) {
    m := map[string][]int{}

    for i := range keyName {
        s := strings.Split(keyTime[i], ":")
        hours, _ := strconv.Atoi(s[0])
        minutes, _ := strconv.Atoi(s[1])
        t := hours * 60 + minutes
        m[keyName[i]] = append(m[keyName[i]], t) 
    }
    for k, v := range m {
        sort.Ints(v)
        n := len(v)
        for i := 0; i < n - 2; i++ {
            if v[i+2] - v[i] <= 60 {
                ans = append(ans, k)
                break
            }
        } 
    }

    sort.Strings(ans)

    return
}