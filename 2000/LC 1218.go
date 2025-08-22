func longestSubsequence(arr []int, difference int) (ans int) {
    f := make(map[int]int)
    for _, x := range arr {
        f[x] = max(f[x], f[x - difference] + 1)
        ans = max(ans, f[x])
    }

    return 
}