func reorderLogFiles(logs []string) []string {
    n := len(logs)
    type pair struct {
        i int
        id, content string
    }
    lettersLog := make([]pair, 0, n)
    for i, log := range logs {
        c := log[len(log) - 1]
        if c >= 'a' && c <= 'z' {
            pos := strings.Index(log, " ")
            lettersLog = append(lettersLog, pair{i, log[:pos], log[pos + 1:]})
        }
    }
    sort.Slice(lettersLog, func(i, j int) bool {
        if lettersLog[i].content == lettersLog[j].content {
            return lettersLog[i].id < lettersLog[j].id
        }
        return lettersLog[i].content < lettersLog[j].content
    })

    ans := make([]string, 0, n)
    for _, v := range lettersLog {
        ans = append(ans, logs[v.i])
    }
    
    for _, log := range logs {
        c := log[len(log) - 1]
        if c < 'a' || c > 'z' {
            ans = append(ans, log)
        }
    }

    return ans
}