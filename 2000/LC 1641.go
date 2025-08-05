const N int = 5
func countVowelStrings(n int) int {
    f := [N]int{}
    for i := range f {
        f[i] = 1
    }

    for i := 1; i < n; i++ {
        for j := 1; j < N; j++ {
            f[j] += f[j-1]
        }
    }

    return f[0] + f[1] + f[2] + f[3] + f[4]
}