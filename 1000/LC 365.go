func canMeasureWater(x int, y int, target int) bool {
    if x + y < target {
        return false
    }

    return target % gcd(x, y) == 0
}

func gcd(x, y int) int {
    for y > 0 {
        x, y = y, x % y
    }

    return x
}