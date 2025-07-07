func canAliceWin(a []string, b []string) bool {
    m, n := len(a), len(b)
    i, j := 0, 0
    turn := 1
    for i < m && j < n {
        if turn == 0 {
            i = solve(i, b[j], a)
        } else {
            j = solve(j, a[i], b)
        }

        turn ^= 1
        fmt.Println(i, j)
    }

    return i < m
}

func solve(i int, s string, a []string) int {
    for i < len(a) {
        x := cmp(a[i], s)
        if x == 1 {
            return i
        } else if x <= 0 {
            i++
        } else {
            break
        }
    }

    return len(a)
}


func cmp(s string, t string) int {
    if s[0] != t[0] {
        return int(s[0]) - int(t[0])
    }
    if s > t {
        return 1
    }
    return 0
}
