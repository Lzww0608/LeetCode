func areaOfMaxDiagonal(dimensions [][]int) int {
    mx, ans := 0, 0
    for _, dimension := range dimensions {
        x, y := dimension[0], dimension[1]
        t := x * x + y * y
        if t > mx {
            mx = t
            ans = x * y
        } else if t == mx {
            ans = max(ans, x * y)
        }
    }

    return ans
}