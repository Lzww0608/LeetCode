func minCost(a []int, b []int) int {
    cnt := make(map[int]int)
    for i := range a {
        cnt[a[i]]++
        cnt[b[i]]++
    }

    for _, v := range cnt {
        if v & 1 == 1 {
            return -1
        }
    }

    for _, x := range a {
        cnt[x] -= 2
    }

    ans := 0
    for _, v := range cnt {
        if v > 0 {
            ans += v / 2
        }
    }

    return ans
}