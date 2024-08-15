func smallestSufficientTeam(req_skills []string, people [][]string) []int {
    m, n := len(req_skills), len(people)
    sid := map[string]int{}
    for i, s := range req_skills {
        sid[s] = i
    }
    mask := make([]int, n)
    for i, v := range people {
        for _, s := range v {
            mask[i] |= 1 << sid[s]
        }
    }

    u := 1 << m
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, u)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if j == 0 {
            return 0
        } else if i < 0 {
            return (1 << n) - 1
        }

        p := &memo[i][j]
        if *p != -1 {
            return *p
        }

        a := dfs(i - 1, j)
        b := dfs(i - 1, j & (^mask[i])) | (1 << i)
        if bits.OnesCount(uint(a)) < bits.OnesCount(uint(b)) {
            *p = a
        } else {
            *p = b
        }

        return *p
    }

    res := dfs(n - 1, u - 1)
    ans := []int{}
    for i := range people {
        if (res >> i) & 1 != 0 {
            ans = append(ans, i)
        }
    }

    return ans
}