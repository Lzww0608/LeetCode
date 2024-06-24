func canCompleteCircuit(gas []int, cost []int) int {
    sum, n := 0, len(gas)
    mx, ans := math.MaxInt32, -1
    for i, x := range cost {
        sum += gas[i] - x
        if sum < mx {
            ans = i
            mx = sum
        }
    }
    if mx > 0 {
        return 0
    }
    if sum < 0 {
        return -1
    }

    return (ans + 1) % n
}