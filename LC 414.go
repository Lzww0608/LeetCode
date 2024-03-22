func thirdMax(nums []int) int {
    a, b, c := math.MinInt, math.MinInt, math.MinInt

    for _, x := range nums {
        if x == a || x == b || x == c {
            continue
        }

        if x > a {
            a, b, c = x, a, b
        } else if x > b {
            b, c = x, b
        } else if x > c {
            c = x
        }
    }

    if c != math.MinInt {
        return c
    }
    return a
}