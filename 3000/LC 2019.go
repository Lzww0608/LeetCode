func scoreOfStudents(s string, answers []int) (ans int) {
    n := len(s)

    correct := 0
    for i := 0; i < n; {
        x := int(s[i] - '0')
        for i += 2; i < n && s[i-1] == '*'; i += 2 {
            x *= int(s[i] - '0')
        }
        correct += x 
    }
    //fmt.Println(correct)
    
    f := make([][]map[int]bool, n)
    for i := range f {
        f[i] = make([]map[int]bool, n)
    }

    var dfs func(int, int) map[int]bool
    dfs = func(l, r int) map[int]bool {
        if l == r {
            return map[int]bool{int(s[l] - '0'): true}
        }

        if f[l][r] != nil {
            return f[l][r]
        }

        res := make(map[int]bool)
        for i := l + 1; i < r; i += 2 {
            for x := range dfs(l, i - 1) {
                for y := range dfs(i + 1, r) {
                    cur := x + y 
                    if s[i] == '*' {
                        cur = x * y 
                    }
                    if cur <= 1000 {
                        res[cur] = true
                    }
                }
            }
        }
        f[l][r] = res 
        return res
    }

    dfs(0, n - 1)

    for _, x := range answers {
        if x == correct {
            ans += 5
        } else if f[0][n-1][x] {
            ans += 2
        }
    }

    return
}