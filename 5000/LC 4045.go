func countGroups(position []int, speed []int, distance int) (ans int) {
    n := len(position)

    pre := 0
    for i := n - 1; i >= 0; i-- {
        if i == n - 1 || speed[i] <= pre && position[i] + distance < position[i + 1] {
            ans++
            pre = speed[i]
        }
    }

    return
}