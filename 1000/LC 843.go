/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
func findSecretWord(word []string, master *Master) {
    n := len(word)
    t := make([][]int, n)
    for i := range t {
        t[i] = make([]int, n)
        t[i][i] = 6
    }

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            cnt := cmp(word[i], word[j])
            t[i][j], t[j][i] = cnt, cnt
        }
    }

    m, idx := n, 0
    for i := range t {
        cnt := [7]int{}
        mx := 0
        for _, j := range t[i] {
            cnt[j]++
            mx = max(mx, cnt[j])
        }

        if mx < m {
            m, idx = mx, i
        }
    }

    s := word[idx]
    ans := master.Guess(s)
    if ans != 6 {
        sub := []string{}
        for _, str := range word {
            if s != str && cmp(s, str) == ans {
                sub = append(sub, str)
            }
        }
        findSecretWord(sub, master)
    }
}

func cmp(a, b string) (ans int) {
    for i := 0; i < 6; i++ {
        if a[i] == b[i] {
            ans++
        }
    }

    return
}