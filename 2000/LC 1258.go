import (
    "strings"
    "slices"
)
func generateSentences(synonyms [][]string, text string) (ans []string) {
    s := strings.Split(text, " ")
    n := len(s)
    k := 0
    id := make(map[string]int)
    rev := make(map[int]string)
    for _, v := range synonyms {
        a, b := v[0], v[1]
        if _, ok := id[a]; !ok {
            id[a] = k
            rev[k] = a
            k++
        }
        if _, ok := id[b]; !ok {
            id[b] = k 
            rev[k] = b
            k++
        }
    }
    m := make([][]bool, k)
    for i := range m {
        m[i] = make([]bool, k)
    }
    for _, v := range synonyms {
        a, b := id[v[0]], id[v[1]]
        m[a][b] = true
        m[b][a] = true
    }

    for i := 0; i < k; i++ {
        for j := i + 1; j < k; j++ {
            if m[i][j] || m[j][i] {
                continue
            }
            for p := 0; p < k; p++ {
                if p != i && p != j {
                    if m[i][p] && m[j][p] {
                        m[i][j] = true
                        m[j][i] = true
                        break
                    }
                }
            }
        }
    }

    var dfs func(int)
    dfs = func(i int) {
        if i >= n {
            ans = append(ans, strings.Join(s, " "))
            return
        }

        dfs(i+1)   
        if _, ok := id[s[i]]; !ok {
            return
        }
        t := id[s[i]]
        for j := 0; j < k; j++ {
            if m[t][j] {
                s[i] = rev[j]
                dfs(i+1)
            }
        }
        return
    }
    dfs(0)

    slices.Sort(ans)
    return
}