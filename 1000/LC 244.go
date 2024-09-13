type WordDistance struct {
    m map[string][]int
}


func Constructor(wordsDict []string) WordDistance {
    m := map[string][]int{}
    for i, s := range wordsDict {
        m[s] = append(m[s], i)
    }

    return WordDistance{m}
}


func (c *WordDistance) Shortest(word1 string, word2 string) int {
    f, g := c.m[word1], c.m[word2]
    m, n := len(f), len(g)
    ans := abs(f[m-1] - g[n-1])
    i, j := 0, 0
    for  i < m && j < n {
        if f[i] < g[j] {
            ans = min(ans, g[j] - f[i])
            i++
        } else if f[i] > g[j] {
            ans = min(ans, f[i] - g[j])
            j++
        } 
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}


/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */