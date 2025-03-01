const N int = 11
func getWordsInLongestSubsequence(words []string, groups []int) (ans []string) {
    s := make([][]int, N)
    for i, t := range words {
        m := len(t)
        s[m] = append(s[m], i)
    }

    seq := []int{}
    for d := 1; d < N; d++ {
        n := len(s[d])
        if n == 0 {
            continue
        }
        f := make([][]int, n)
        f[0] = []int{s[d][0]}
        if len(f[0]) > len(seq) {
            seq = f[0]
        }
        for j := 1; j < n; j++ {
            mx := -1
            for k := j - 1; k >= 0; k-- {
                if check(words[s[d][k]], words[s[d][j]]) && groups[s[d][k]] != groups[s[d][j]] {
                    if mx == -1 || len(f[k]) > len(f[mx]) {
                        mx = k
                    }
                }
            }
            if mx != -1 {
                f[j] = append(append([]int(nil), f[mx]...), s[d][j])
            } else {
                f[j] = []int{s[d][j]}
            }
            
            if len(f[j]) > len(seq) {
                seq = f[j]
            }
        }
    }
    for _, x := range seq {
        ans = append(ans, words[x])
    }

    return
}


func check(s, t string) bool {
    cnt := 0
    for i := range s {
        if s[i] != t[i] {
            cnt++
        }
    }

    return cnt == 1
} 