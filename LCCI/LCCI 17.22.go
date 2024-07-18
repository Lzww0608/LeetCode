func findLadders(beginWord string, endWord string, wordList []string) (ans []string) {
    check := func(src, target string) bool {
        if len(src) != len(target) {
            return false
        }
        cnt := 0
        for i := range src {
            if src[i] != target[i] {
                cnt++
            }
        }

        return cnt == 1
    }

    vis := make([]bool, len(wordList))
    var dfs func(int) bool
    dfs = func(idx int) bool {
        if wordList[idx] == endWord {
            ans = append(ans, endWord)
            return true
        } else if vis[idx] {
            return false
        }

        ans = append(ans, wordList[idx])
        vis[idx] = true
        for i, s := range wordList {
            if check(s, wordList[idx]) {
                if dfs(i) {
                    return true
                }
            }
        }
        ans = ans[:len(ans)-1]
        return false
    }

    ans = append(ans, beginWord)
    for i, s := range wordList {
        if check(s, beginWord) {
            if dfs(i) {
                return 
            }
        }
    }

    return nil
}