func wateringPlants(plants []int, capacity int) (ans int) {
    k := capacity
    for i, x := range plants {
        if k >= x {
            k -= x
            ans++
        } else {
            k = capacity - x
            ans += 1 + i * 2
        }
    }

    return
}
