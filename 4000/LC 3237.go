func simulationResult(windows []int, queries []int) []int {
    pos := make(map[int]int)
    for i, window := range windows {
        pos[window] = i
    }

    cur := -1
    for _, q := range queries {
        pos[q] = cur 
        cur--
    }

    sort.Slice(windows, func(i, j int) bool {
        return pos[windows[i]] < pos[windows[j]]
    })

    return windows
}