func tictactoe(moves [][]int) string {
    a := [3][3]byte{}
    check := func() byte {
        for i := 0; i < 3; i++ {
            f := true
            for j := 1; j < 3; j++ {
                if a[i][j] != a[i][j-1] {
                    f = false
                    break
                }
            }
            if f && a[i][0] != 0 {
                return a[i][0]
            }
        }

        for j := 0; j < 3; j++ {
            f := true
            for i := 1; i < 3; i++ {
                if a[i][j] != a[i-1][j] {
                    f = false
                    break
                }
            }
            if f && a[0][j] != 0 {
                return a[0][j]
            }
        }

        if a[0][0] == a[1][1] && a[1][1] == a[2][2] && a[0][0] != 0 {
            return a[0][0]
        }
        if a[0][2] == a[1][1] && a[1][1] == a[2][0] && a[0][2] != 0 {
            return a[0][2]
        }

        return ' '
    }

    for k, v := range moves {
        i, j := v[0], v[1]
        if k & 1 == 0 {
            a[i][j] = 'X'
        } else {
            a[i][j] = 'O'
        }
        c := check()
        if c == 'X' {
            return "A"
        } else if c == 'O' {
            return "B"
        }
    }

    if len(moves) == 9 {
        return "Draw"
    }
    return "Pending"
}