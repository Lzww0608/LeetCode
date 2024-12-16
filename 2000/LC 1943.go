const N int = 1_000_05
func splitPainting(segments [][]int) (ans [][]int64) {
    sort.Slice(segments, func(i, j int) bool {
        return segments[i][0] < segments[j][0] || (segments[i][0] == segments[j][0] && segments[i][1] < segments[j][1])
    })

    dif := make([]int, N)
    vis := make(map[int]bool)
    for _, v := range segments {
        dif[v[0]] += v[2]
        dif[v[1]] -= v[2]
        vis[v[0]] = true
        vis[v[1]] = true
    }

    cur, l := 0, 0
    for r := 0; r < N; r++ {
        if !vis[r]{
            continue
        }
        if cur > 0 {
            ans = append(ans, []int64{int64(l), int64(r), int64(cur)}) 
        }
        cur += dif[r]
        l = r
    }

    return
}