func confusingNumberII(n int) int {
    s := strconv.Itoa(n)
    m := map[byte]byte {
        '0': '0',
        '1': '1',
        '6': '9',
        '8': '8',
        '9': '6',
    }

    var dfs func(int, string, bool) int
    dfs = func(i int, cur string, isLimit bool) (ans int) {
        if i == len(s) {
            a, _ := strconv.Atoi(cur)
            cur = strconv.Itoa(a)
            t := make([]byte, len(cur))
            for i := range t {
                if _, ok := m[cur[i]]; !ok {
                    return 0
                }
                t[len(t)-i-1] = m[cur[i]]
            }
            if cur != string(t) {
                //fmt.Println(cur, string(t))
                return 1
            }
            return 0
        }

        low, up := 0, 9 
        if isLimit {
            up = int(s[i] - '0')
        }

        for d := low; d <= up; d++ {
            c := byte('0' + d)
            if _, ok := m[c]; !ok {
                continue
            }
            ans += dfs(i + 1, cur + string(c), isLimit && d == up)
        }

        return
    }

    return dfs(0, "", true)
}

