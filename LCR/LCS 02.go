const N int = 1_001
func halfQuestions(questions []int) (ans int) {
    cnt := make([]int, N)
    n := len(questions) / 2
    for _, x := range questions {
        cnt[x]++
    }

    sort.Ints(cnt)
    sum := 0

    for i := N - 1; i >= 0 && sum < n; i-- {
        sum += cnt[i]
        ans++
    }

    return
}