var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

type move struct {
    x, y int
    dx, dy int
    step int
}

const N int = 8

func isValid(a, b move) bool {
    a_i, a_j := a.x, a.y
    b_i, b_j := b.x, b.y
    for i := range max(a.step, b.step) {
        if i < a.step {
            a_i += a.dx
            a_j += a.dy
        }
        if i < b.step {
            b_i += b.dx
            b_j += b.dy
        }
        if a_i == b_i && a_j == b_j {
            return false
        } 
    }
    return true
}

func generate(i, j int, dirs [][]int) []move {
    ans := []move{{i, j, 0, 0, 0}}
    for _, dir := range dirs {
        x, y := i + dir[0], j + dir[1]
        for t := 1; x > 0 && x <= N && y > 0 && y <= N; t++ {
            ans = append(ans, move{i, j, dir[0], dir[1], t})
            x += dir[0]
            y += dir[1]
        }
    }

    return ans
}

func countCombinations(pieces []string, positions [][]int) (ans int) {
    n := len(pieces)
    moves := make([][]move, n)
    m := map[string][]int{
        "rook": {0, 4},
        "queen": {0, 8},
        "bishop": {4, 8},
    }
    for i, pos := range positions {
        moves[i] = generate(pos[0], pos[1], dirs[m[pieces[i]][0]:m[pieces[i]][1]])
    }

    path := make([]move, n)
    var dfs func(int)
    dfs = func(i int) {
        if i == n {
            ans++
            return
        }
        next:
        for _, a := range moves[i] {
            for _, b := range path[:i] {
                if !isValid(a, b) {
                    continue next
                }
            }
            path[i] = a
            dfs(i + 1)
        }
    }
    dfs(0)
    return
}