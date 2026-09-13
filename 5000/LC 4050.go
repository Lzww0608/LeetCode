const N int = 100_001
var F[N]int

func init() {
    for i := 0; i < N - 1; i++ {
        t := i
        for j := 1; t + j < N; j++ {
            t += j
            if F[t] == 0 || F[t] > F[i] + j + 1 {
                F[t] = F[i] + j + 1
            }
        } 
    }
}

func minDays(n int) int {
    return F[n] - 1 
}