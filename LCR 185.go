func statisticsProbability(num int) []float64 {
    const N int = 6

    ans := make([]float64, N)
    for i := range ans {
        ans[i] = float64(1) / float64(N)
    }

    for i := 2; i <= num; i++ {
        tmp := make([]float64, i * (N - 1) + 1)
        for j := range ans {
            for k := 0; k < N; k++ {
                tmp[j+k] += ans[j] / float64(N)
            }
        }
        ans = tmp
    }

    return ans
}