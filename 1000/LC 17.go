var alp []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
func letterCombinations(digits string) (ans []string) {
    n := len(digits)
    if n == 0 {
        return
    }
    path := make([]rune, n)
    var dfs func(int)
    dfs = func(id int) {
        if id == n {
            ans = append(ans, string(path))
            return
        }
        for _, c := range alp[int(digits[id]-'0')] {
            path[id] = c
            dfs(id+1)
        }
    }
    dfs(0)
    return 
}
