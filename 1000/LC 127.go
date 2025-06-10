const N int = 26
func ladderLength(beginWord string, endWord string, wordList []string) int {
    m := make(map[string]bool)
    for _, s := range wordList {
        m[s] = true
    }

    if !m[endWord] {
        return 0
    }

    q := []string{beginWord}
    d := 1
    for len(q) > 0 {
        for sz := len(q); sz > 0; sz-- {
            if q[0] == endWord {
                return d
            }
            cur := []byte(q[0])
            q = q[1:]

            for i := range cur {
                origin := cur[i]
                for j := range N {
                    if j != int(origin - 'a') {
                        cur[i] = byte('a' + j)
                        if m[string(cur)] {
                            m[string(cur)] = false
                            q = append(q, string(cur))
                        }
                        cur[i] = origin
                    }
                } 
            }
        }
        d++
    }

    return 0
}