func findClosedNumbers(num int) []int {
    if num == math.MaxInt32 {
        return []int{-1, -1}
    } else if num == 1 {
        return []int{2, -1}
    }

    return []int{Gosper(num), ^Gosper(^num)}
}

func Gosper(x int) int {
    c := x & -x
    r := x + c

    return (r ^ x) / c >> 2 | r
}