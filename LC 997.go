func findJudge(n int, trust [][]int) int {
    in, out := make([]int, n), make([]int, n)
    for _, e := range trust {
        a, b := e[0] - 1, e[1] - 1
        out[a]++
        in[b]++
    }
    for i := range in {
        if in[i] == n - 1 && out[i] == 0 {
            return i + 1
        }
    }

    return -1
}