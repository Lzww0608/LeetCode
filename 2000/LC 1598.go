func minOperations(logs []string) (ans int) {
    for _, s := range logs {
        if s == "../" {
            ans = max(ans - 1, 0)
        } else if s != "./" {
            ans++
        }
    }

    return
}