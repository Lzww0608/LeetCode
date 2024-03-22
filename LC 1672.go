func maximumWealth(accounts [][]int) (ans int) {
    for i := range accounts {
        wealth := 0
        for _, x := range accounts[i] {
            wealth += x
        }
        ans = max(ans, wealth)
    }

    return
}