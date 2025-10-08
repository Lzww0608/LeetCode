func findCommonResponse(responses [][]string) (ans string) {
    cnt := make(map[string]int)
    for _, response := range responses {
        m := make(map[string]bool)
        for _, s := range response {
            m[s] = true
        }

        for k := range m {
            cnt[k]++
        }
    }

    for k, v := range cnt {
        if v > cnt[ans] || v == cnt[ans] && k < ans {
            ans = k
        }
    }

    return 
}