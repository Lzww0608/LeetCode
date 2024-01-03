func reachNumber(target int) int {
    d, ans := 0, 0
    target = abs(target)
    for true {
        ans++
        d += ans
        if d >= target && (d - target) % 2 == 0 {
            break
        }
    }
    return ans
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}