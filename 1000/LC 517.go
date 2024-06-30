func findMinMoves(machines []int) (ans int) {
    tot, n := 0, len(machines)
    for _, x := range machines {
        tot += x
    }

    if tot % n != 0 {
        return -1
    }

    sum, avg := 0, tot / n
    for _, x := range machines {
        x -= avg
        sum += x
        ans = max(ans, sum, -sum, x)
    }

    return
}