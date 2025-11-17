func minArrivalsToDiscard(arrivals []int, w int, m int) (ans int) {
    cnt := make(map[int]int)
    dis := make([]bool, len(arrivals))
    for i, x := range arrivals {
        if cnt[x] == m {
            ans++
            dis[i] = true
        } else {
            cnt[x]++
        }
        
        if i >= w - 1 && !dis[i - w + 1] {
            cnt[arrivals[i - w + 1]]--
        }
    }

    return
}