func winnerOfGame(colors string) bool {
    n := len(colors)
    cnt := 0

    for i := 1; i < n - 1; i++ {
        if colors[i] == colors[i - 1] && colors[i] == colors[i + 1] {
            if colors[i] == 'A' {
                cnt++
            } else {
                cnt--
            }
        }
    }

    return cnt > 0
}