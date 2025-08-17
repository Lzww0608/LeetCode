func findPattern(board [][]int, pattern []string) []int {
    m1, n1 := len(board), len(board[0])
    m2, n2 := len(pattern), len(pattern[0])
    if m2 > m1 || n2 > n1 {
        return []int{-1, -1}
    }

    for i := 0; i + m2 <= m1; i++ {
        next:
        for j := 0; j + n2 <= n1; j++ {
            letter2Digit := make(map[byte]int)
            digit2Letter := make(map[int]byte)
            for dx := 0; dx < m2; dx++ {
                for dy := 0; dy < n2; dy++ {
                    c := pattern[dx][dy]
                    t := board[i + dx][j + dy]
                    if c >= '0' && c <= '9' {
                        if int(c - '0') != t {
                            continue next
                        }
                    } else {
                        v1, ok1 := letter2Digit[c]
                        v2, ok2 := digit2Letter[t] 
                        if ok1 && v1 != t {
                            continue next
                        } else if ok2 && v2 != c {
                            continue next
                        }
                        letter2Digit[c] = t
                        digit2Letter[t] = c
                    }
                }
            }

            return []int{i, j}
        }
    }

    return []int{-1, -1}
}