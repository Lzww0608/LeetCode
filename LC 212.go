type TrieNode struct {
    children [26]*TrieNode
    word     string
}

func init() {
    debug.SetGCPercent(-1)
}

func findWords(board [][]byte, words []string) []string {
    root := &TrieNode{}
    for _, word := range words {
        node := root
        for _, char := range word {
            if node.children[char-'a'] == nil {
                node.children[char-'a'] = &TrieNode{}
            }
            node = node.children[char-'a']
        }
        node.word = word
    }

    m, n := len(board), len(board[0])
    result := []string{}
    var dfs func(node *TrieNode, i, j int)
    dfs = func(node *TrieNode, i, j int) {
        char := board[i][j]
        if char == '#' || node.children[char-'a'] == nil {
            return
        }
        node = node.children[char-'a']
        if node.word != "" {
            result = append(result, node.word)
            node.word = "" 
        }

        board[i][j] = '#'
        dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n {
                dfs(node, x, y)
            }
        }
        board[i][j] = char
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(root, i, j)
        }
    }

    return result
}
