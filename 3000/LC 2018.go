func placeWordInCrossword(board [][]byte, word string) bool {
    m, n, k := len(board), len(board[0]), len(word)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == '#' {
                continue
            }

            p, ok1, ok2 := j, true, true
            for p < n && board[i][p] != '#' {
                if p - j >= k || ok1 && board[i][p] != ' ' && board[i][p] != word[p-j] {
                    ok1 = false
                }

                if p - j >= k || ok2 && board[i][p] != ' ' && board[i][p] != word[k-1-(p-j)] {
                    ok2 = false
                }
                p++
            }
            if p - j == k && (ok1 || ok2) {
                return true
            }
            
            j = p
        }
    }

    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            if board[i][j] == '#' {
                continue
            }

            p, ok1, ok2 := i, true, true
            for p < m && board[p][j] != '#' {
                if p - i >= k || ok1 && board[p][j] != ' ' && board[p][j] != word[p-i] {
                    ok1 = false
                }

                if p - i >= k || ok2 && board[p][j] != ' ' && board[p][j] != word[k-1-(p-i)] {
                    ok2 = false
                }
                p++
            }
            if p - i == k && (ok1 || ok2) {
                return true
            }
            
            i = p
        }
    }

    return false
}