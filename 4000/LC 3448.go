const N int = 10
func countSubstrings(s string) int64 {
    ans := 0

    mod := [N][N]int{}
    for k := range s {
        x := int(s[k] - '0')
        for i := 1; i < N; i++ {
            tmp := [N]int{}
            tmp[x % i] = 1
            for j := 0; j < i; j++ {
                tmp[(j * 10 + x) % i] += mod[i][j]
            }
            mod[i] = tmp
        }

        ans += mod[x][0]
    }

    return int64(ans)
}