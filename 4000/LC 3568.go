func minMoves(classroom []string, energy int) int {
    dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    m, n := len(classroom), len(classroom[0])
    id := make(map[int]int)
    k := 0

    type pair struct {
        x, e, c, d int 
    }func minMoves(classroom []string, energy int) int {
    dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    m, n := len(classroom), len(classroom[0])
    id := make(map[int]int)
    k := 0

    type pair struct {
        x, e, c, d int 
    }
    q := []pair{}
    memo := make([][]bool, m * n)
    for i := range memo {
        memo[i] = make([]bool, (energy << 11) + (1 << 10) + 1)
    }

    for i := range m {
        for j := range n {
            if classroom[i][j] == 'L' {
                id[i * n + j] = k 
                k++
            } else if classroom[i][j] == 'S' {
                q = append(q, pair{i * n + j, energy, 0, 0})
                memo[i * n + j][energy << 11] = true 
            }
        }
    }
    if k == 0 {
        return 0
    }

    mask := (1 << k) - 1
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        i, j := cur.x / n, cur.x % n
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || classroom[x][y] == 'X' {
                continue
            }
            e, c := cur.e - 1, cur.c
            if classroom[x][y] == 'R' {
                e = energy
            } else if classroom[x][y] == 'L' {
                k = id[x * n + y]
                if c & (1 << k) == 0 {
                    c |= 1 << k
                }
            }

            if c == mask {
                return cur.d + 1
            } else if e == 0 {
                continue
            }

            p := x * n + y
            if !memo[p][(e << 11) + c] {
                memo[p][(e << 11) + c] = true 
                q = append(q, pair{p, e, c, cur.d + 1})
            }
        }
    }
    
    return -1
}
    q := []pair{}

    for i := range m {
        for j := range n {
            if classroom[i][j] == 'L' {
                id[i * n + j] = k 
                k++
            } else if classroom[i][j] == 'S' {
                q = append(q, pair{i * n + j, energy, 0, 0})
            }
        }
    }

    mask := (1 << k) - 1
    memo := make([][]bool, m * n)
    for i := range memo {
        memo[i] = make([]bool, (energy << 11) + (1 << 10) + 1)
    }
    memo[0][energy << 11] = true 
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        i, j := cur.x / n, cur.x % n
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || classroom[x][y] == 'X' {
                continue
            }
            e, c := cur.e - 1, cur.c
            if classroom[x][y] == 'R' {
                e = energy
            } else if classroom[x][y] == 'L' {
                k = id[x * n + y]
                if c & (1 << k) == 0 {
                    c |= 1 << k
                }
            }

            if c == mask {
                return cur.d + 1
            } else if e == 0 {
                continue
            }

            p := x * n + y
            if !memo[p][(e << 11) + c] {
                memo[p][(e << 11) + c] = true 
                q = append(q, pair{p, e, c, cur.d + 1})
            }
        }
    }
    
    return -1
}