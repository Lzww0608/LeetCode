func findHighAccessEmployees(access_times [][]string) (ans []string) {
    cnt := make(map[string][]int)
    for _, v := range access_times {
        x, _ := strconv.Atoi(v[1][2:])
        x += int(v[1][0] - '0') * 600 + int(v[1][1] - '0') * 60
        if _, ok := cnt[v[0]]; !ok {
            cnt[v[0]] = make([]int, 0)
        }
        cnt[v[0]] = append(cnt[v[0]], x)
    }

    for k, v := range cnt {
        sort.Ints(v)
        for i := 0; i + 2 < len(v); i++ {
            if v[i + 2] - v[i] < 60 {
                ans = append(ans, k)
                break
            }
        }
    }

    return
}