func numOfBurgers(a int, b int) []int {
    // x + y == b 
    // x * 4 + y * 2 == a 
    // x == b - y
    // 4b - 4y + 2y == a 
    // 4b - a == 2y 
    if (b * 4 - a) & 1 == 1 {
        return nil
    }

    y := (b * 4 - a) / 2
    x := b - y
    if x < 0 || y < 0 {
        return nil
    }
    return []int{x, y}
}