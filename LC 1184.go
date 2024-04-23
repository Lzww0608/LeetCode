func distanceBetweenBusStops(distance []int, start int, destination int) int {
    a, b := 0, 0
    n := len(distance)
    for i := start; i != destination; i = (i + 1) % n {
        a += distance[i]
    }
    for i := start; i != destination; i = (i - 1 + n) % n {
        b += distance[(i - 1 + n) % n]
    }

    return min(a, b)
}