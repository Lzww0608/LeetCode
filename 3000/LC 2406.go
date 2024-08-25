func minGroups(intervals [][]int) (ans int) {
    cnt := make([]int, 1_000_005)
    for _, v := range intervals {
        cnt[v[0]]++
        cnt[v[1]+1]--
    }

    cur := 0
    for _, x := range cnt {
        if x == 0 {
            continue
        }
        cur += x
        ans = max(ans, cur)
    }

    return 
}