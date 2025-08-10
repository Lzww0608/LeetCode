const N int = 26
func sameEndSubstringCount(s string, queries [][]int) []int {
    n := len(s)
    pre := make([][N]int, n + 1)
    p := [N]int{}

    for i := range s {
        p[int(s[i] - 'a')]++
        pre[i + 1] = p
    }

    ans := make([]int, len(queries))
    for i, query := range queries {
        l, r := query[0], query[1]
        for j := 0; j < N; j++ {
            x := pre[r + 1][j] - pre[l][j]
            ans[i] += x * (x + 1) / 2 
        }
    }

    return ans
}