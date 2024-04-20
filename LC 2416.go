func sumPrefixScores(words []string) []int {
    type node struct {
        son [26]*node
        score int
    }

    root := &node{}
    for _, s := range words {
        cur := root
        for _, c := range s {
            c -= 'a'
            if cur.son[c] == nil {
                cur.son[c] = &node{}
            } 
            cur = cur.son[c]
            cur.score++
        }
    }

    ans := make([]int, len(words))
    for i, s := range words {
        cur := root
        for _, c := range s {
            c -= 'a'
            cur = cur.son[c]
            ans[i] += cur.score
        }
    }

    return ans
}