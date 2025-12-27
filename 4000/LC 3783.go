func mirrorDistance(n int) int {
    return abs(n - reverse(n))
}

func reverse(n int) (ans int) {
    for n > 0 {
        ans = ans * 10 + n % 10
        n /= 10 
    }

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}