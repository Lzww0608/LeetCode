func minMaxWaitingTime(demand []int, fuel []int) int {
    n := len(demand)
    memo := make(map[int][2]int)

    var dfs func(i, wait0, wait1, fuel0, fuel1 int) (int, int) 
    dfs = func(i, wait0, wait1, fuel0, fuel1 int) (int, int)   {
        if i == n {
            return 0, 0
        }

        mask := i << 24 | wait0 << 18 | wait1 << 12 | fuel0 << 6 | fuel1
        if v, ok := memo[mask]; ok {
            return v[0], v[1]
        }

        mx, best := 0, 0
        d := demand[i]
        if d <= fuel0 {
            num, time := dfs(i + 1, d, max(wait1 - wait0, 0), fuel0 - d, fuel1)
            mx = num + 1
            best = max(time, wait0)
        }

        if d <= fuel1 {
            num, time := dfs(i + 1, max(wait0 - wait1, 0), d, fuel0, fuel1 - d)
            num++
            time = max(time, wait1)
            if num > mx || num == mx && time < best {
                best = time 
                mx = num
            }
        }
        memo[mask] = [2]int{mx, best}

        return mx, best
    }

    u, v := dfs(0, 0, 0, fuel[0], fuel[1])
    if u == 0 {
        return -1
    }
    return v
}