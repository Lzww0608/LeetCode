func findColoredRoot(edges [][]int, colors string) int {
    col := map[byte]int {
        'R': 0,
        'B': 1,
        'W': 2,
    }
    n := len(colors)
    g := make([][]int, n)
    for _, v := range edges {
        g[v[0]] = append(g[v[0]], v[1])
        g[v[1]] = append(g[v[1]], v[0])
    }

    ans := -1
    for x, v := range g {
        if len(v) > 3 {
            return -1
        }

        cnt := [3]int{}
        for _, y := range v {
            cnt[col[colors[y]]]++
        }

        c := col[colors[x]]
        if cnt[c] > 0 || cnt[(c - 1 + 3) % 3] > 1 {
            return -1
        }

        if len(v) < 3 && cnt[(c - 1 + 3) % 3] == 0 {
            ans = x
        }
    }

    return ans
}