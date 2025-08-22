func isWinner(player1 []int, player2 []int) int {
    a, b := 0, 0
    n := len(player1)
    for i := 0; i < n; i++ {
        if i - 1 >= 0 && player1[i - 1] == 10 || i - 2 >= 0 && player1[i - 2] == 10 {
            a += player1[i] * 2  
        } else {
            a += player1[i]
        }

        if i - 1 >= 0 && player2[i - 1] == 10 || i - 2 >= 0 && player2[i - 2] == 10 {
            b += player2[i] * 2  
        } else {
            b += player2[i]
        }
    }

    if a > b {
        return 1
    } else if a < b {
        return 2
    }

    return 0
}