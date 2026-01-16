func recoverOrder(order []int, friends []int) []int {
    n := len(order)
    vis := make([]bool, n)
    for _, x := range friends {
        vis[x - 1] = true
    }

    ans := make([]int, 0, len(friends))
    for _, x := range order {
        if vis[x - 1] {
            ans = append(ans, x)
        }
    }

    return ans
}