const N int = 1_001
func numRabbits(answers []int) (ans int) {
    cnt := make([]int, N)
    for _, x := range answers {
        cnt[x]++
    }

    for i := 0; i < N; i++ {
        t := (cnt[i] + i) / (i + 1)
        ans += t * (i + 1)
    }

    return
}