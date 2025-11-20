const N int = 8
const M int = 26
func minimumPushes(word string) (ans int) {
    cnt := make([]int, M)
    for _, c := range word {
        x := int(c - 'a')
        cnt[x]++
    }

    sort.Ints(cnt)
    s, t := N, 1
    for i := M - 1; i >= 0; i-- {
        if cnt[i] == 0 {
            break
        }
        ans += cnt[i] * t 
        s--
        if s == 0 {
            s, t = N, t + 1
        }
    }

    return 
}