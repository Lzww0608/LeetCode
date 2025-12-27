func getMinimumTime(time []int, fruits [][]int, limit int) (ans int) {
    for _, v := range fruits {
        t, n := v[0],  v[1]
        ans += (n + limit - 1) / limit * time[t]
    }

    return 
}