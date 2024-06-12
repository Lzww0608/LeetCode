
type pair struct {
    idx int
    str string
}

type WordFilter struct {
    pre, suf map[string][]pair
}

func Constructor(words []string) WordFilter {
    pre, suf := make(map[string][]pair), make(map[string][]pair)
    for j, s := range words {
        for i := 1; i <= len(s); i++ {
            pre[s[:i]] = append(pre[s[:i]], pair{j, s})
        }
        for i := 0; i < len(s); i++ {
            suf[s[i:]] = append(suf[s[i:]], pair{j, s})
        }
    }
    return WordFilter{pre, suf}
}

func (wf *WordFilter) F(pref string, suff string) int {
    preMatches, sufMatches := wf.pre[pref], wf.suf[suff]
    i, j := len(preMatches)-1, len(sufMatches)-1

    for i >= 0 && j >= 0 {
        if preMatches[i].idx == sufMatches[j].idx {
            return preMatches[i].idx
        }
        if preMatches[i].idx > sufMatches[j].idx {
            i--
        } else {
            j--
        }
    }
    return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
