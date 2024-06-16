func minFlips(target string) (ans int) {
    n := len(target)
    for i := 1; i < n; i++ {
        if target[i] != target[i-1] {
            ans++
        }
    }
    if target[0] == '1' {
        ans++
    }
    return ans
}