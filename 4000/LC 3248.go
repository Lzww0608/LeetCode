func finalPositionOfSnake(n int, commands []string) int {
    x, y := 0, 0
    for _,  s := range commands {
        switch s {
            case "UP":
                x, y = x - 1, y
            case "DOWN":
                x, y = x + 1, y 
            case "LEFT":
                x, y = x, y - 1
            case "RIGHT":
                x, y = x, y + 1
        }
    }

    return x * n + y
}