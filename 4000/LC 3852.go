func minDistinctFreqPair(nums []int) []int {
    cnt := make(map[int]int)
    mn := math.MaxInt32
    for _, x := range nums {
        cnt[x]++
        mn = min(mn, x)
    }

    y := math.MaxInt32
    for k, v := range cnt {
        if k == mn {
            continue
        }
        if v != cnt[mn] {
            y = min(y, k)
        }
    }

    if y == math.MaxInt32 {
        return []int{-1, -1}
    }

    return []int{mn, y}
}