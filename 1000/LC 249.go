func groupStrings(strings []string) (ans [][]string) {
    m := map[string][]int{}

    for i, s := range strings {
        n := len(s)
        a := make([]byte, n)
        a[0] = 'a'
        for i := 1; i < n; i++ {
            t := int(s[i] - s[i-1] + 26) % 26
            a[i] = byte('a' + t)
        }
        m[string(a)] = append(m[string(a)], i) 
    }

    for _, v := range m {
        tmp := []string{}
        for _, i := range v {
            tmp = append(tmp, strings[i])
        }
        ans = append(ans, tmp)
    }

    return
}