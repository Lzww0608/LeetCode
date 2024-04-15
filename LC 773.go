func slidingPuzzle(board [][]int) int {

    neightbor := [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
    target := "123450"
    builder := strings.Builder{}
    for _, row := range board {
        for _, x := range row {
            builder.WriteString(strconv.Itoa(x))
        } 
    }
…        k++
    }

    return -1
}