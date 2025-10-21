func longestCommonPrefix(s []string) []int {
    n := len(s)
    ans := make([]int, n)
    a, b, c := 0, 0, 0

    for i := 0; i < n - 1; i++ {
        x := solve(s[i], s[i + 1])
        if x >= a {
            a, b, c = x, a, b
        } else if x >= b {
            b, c = x, b
        } else if x > c {
            c = x
        }
    }

    for i := range n {
        var x, y, z int 
        if i > 0 {
            x = solve(s[i], s[i - 1])
            if i < n - 1 {
                z = solve(s[i - 1], s[i + 1])
            }
        } 
        if i < n - 1 {
            y = solve(s[i], s[i + 1])
        }

        if x < y {
            x, y = y, x 
        }
        if x == a {
            if y == b {
                ans[i] = c
            } else {
                ans[i] = b
            }
        } else {
            ans[i] = a
        }
        ans[i] = max(ans[i], z)
    }

    return ans
}

func solve(s, t string) int {
    for i := 0; i < min(len(s), len(t)); i++ {
        if s[i] != t[i] {
            return i
        }
    }

    return min(len(t), len(s))
}