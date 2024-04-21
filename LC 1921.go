func eliminateMaximum(dist []int, speed []int) int {
    n := len(dist)
    cnt := make([]int, n)

    for i := range dist {
        time := (dist[i] + speed[i] - 1) / speed[i] 
        if time <= n {
            cnt[time-1]++
        }
    }

    kill := 0
    for i := range cnt {
        kill++
        kill -= cnt[i]
        if kill < 0 {
            return i + 1
        }
    }
    return n
}