func elevatorRequests(n int, requests []int) (ans int) {
    y := 0
    for _, x := range requests {
        ans += abs(x - y)
        y = x
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}