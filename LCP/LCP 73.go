func adventureCamp(expeditions []string) int {
    start := make(map[string]bool)
    s := strings.Split(expeditions[0], "->")
    for _, t := range s {
        start[t] = true
    }

    mx, ans := 0, -1
    vis := make(map[string]bool)
    for i, v := range expeditions[1:] {
        if len(v) == 0 {
            continue
        }
        clear(vis)
        for _, t := range strings.Split(v, "->") {
            if !start[t] {
                vis[t] = true
            }
        }

        if mx < len(vis) {
            ans = i + 1
            mx = len(vis)
        }
        for t := range vis {
            start[t] = true
        }
    }

    return ans
}