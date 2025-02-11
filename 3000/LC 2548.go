
func maxPrice(items [][]int, capacity int) (ans float64) {
    sort.Slice(items, func(i, j int) bool {
        return items[i][0] * items[j][1] > items[j][0] * items[i][1]
    })    

    for _, v := range items {
        t := min(v[1], capacity)
        ans += float64(t) * float64(v[0]) / float64(v[1])
        capacity -= t
        if (capacity == 0) {
            break
        }
    }

    if (capacity != 0) {
        return -1
    }
    return
}