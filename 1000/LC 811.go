func subdomainVisits(cpdomains []string) (ans []string) {
    cnt := make(map[string]int)
    for _, cpdomain := range cpdomains {
        s := strings.Split(cpdomain, " ")
        rep, _ := strconv.Atoi(s[0])
        domain := s[1]
        for {
            cnt[domain] += rep
            if idx := strings.IndexByte(domain, '.'); idx == -1 {
                break
            } else {
                domain = domain[idx+1:] // 下一级子域
        }
        }
    }

    for k, v := range cnt {
        s := strconv.Itoa(v) + " " + k
        ans = append(ans, s)
    }

    return
}