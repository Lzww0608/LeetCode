func kItemsWithMaximumSum(a int, b int, c int, k int) (ans int) {
    ans += min(a, k)
    k -= min(a, k)

    k -= min(b, k)
    
    ans -= min(c, k)

    return
}