func findGCD(nums []int) int {
    mx, mn := nums[0], nums[0]
    for _, x := range nums {
        if mx < x {
            mx = x
        }
        if mn > x {
            mn = x
        }
    }

    return gcd(mx, mn)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}