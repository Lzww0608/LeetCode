const N int = 26
func minimumPushes(word string) (ans int) {
    cnt := make([]int, N)
    for _, c := range word {
        cnt[int(c - 'a')]++
    }
    sort.Ints(cnt)

    sum, pos := 0, 1
    for i := N - 1; i >= 0; i-- {
        if cnt[i] == 0 {
            break
        }

        ans += pos * cnt[i]
        sum++
        if sum == 8 {
            sum = 0
            pos++
        }
    }

    return
}