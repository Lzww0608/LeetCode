func letterCasePermutation(s string) (ans []string) {
    a := []byte(s)
    for i := range a {
        if a[i] >= 'A' && a[i] <= 'Z' {
            a[i] += 32
        }
    }

    var dfs func(int) 
    dfs = func(i int) {
        if i == len(a) {
            ans = append(ans, string(a))
            return
        }
        dfs(i + 1)
        if a[i] >= 'a' && a[i] <= 'z' {
            a[i] -= 32 
            dfs(i + 1)
            a[i] += 32
        }
    }

    dfs(0)

    return
}