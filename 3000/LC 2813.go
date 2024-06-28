func findMaximumElegance(items [][]int, k int) (ans int64) {
    sort.Slice(items, func(i, j int) bool {return items[i][0] > items[j][0]})
    sum := 0
    vis := map[int]bool{}
    st := []int{}

    for i, v := range items {
        pro, cat := v[0], v[1]
        if i < k {
            sum += pro
            if !vis[cat] {
                vis[cat] = true
            } else {
                st = append(st, pro)
            }
        } else if len(st) > 0 && !vis[cat] {
            vis[cat] = true
            sum += pro - st[len(st)-1]
            st = st[:len(st)-1]
        }
        ans = max(ans, int64(sum + len(vis) * len(vis)))
    }

    return ans
}
