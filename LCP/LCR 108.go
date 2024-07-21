func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := make(map[string]bool)
    for _, s := range wordList {
        wordSet[s] = true
    }

    if !wordSet[endWord] {
        return 0
    }

    q := []string{beginWord}
    step := 1
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            cur := q[0]
            q = q[1:]
            if cur == endWord {
                return step
            } 

            for i := range cur {
                for j := 0; j < 26; j++ {
                    c := byte('a' + j)
                    if cur[i] != c {
                        s := cur[:i] + string(c) + cur[i+1:]
                        if wordSet[s] {
                            q = append(q, s)
                            wordSet[s] = false
                        }
                    }
                }
            }
        }
        step++
    }

    return 0
}