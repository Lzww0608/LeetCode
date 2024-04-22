func numMovesStones(a int, b int, c int) []int {
    sum := a + b + c
    a, c = min(a, b, c), max(a, b, c)
    b = sum - a - c
    mn := 2
    if c - b == 1 && b - a == 1 {
        mn = 0
    } else if c - b <= 2 || b - a <= 2{
        mn = 1
    }
    return []int{mn , max(0, c - a - 2)}
}