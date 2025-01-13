func minimumSplits(nums []int) (ans int) {
    g := 1
    for _, x := range nums {
        g = gcd(g, x)
        if g == 1 {
            ans++
            g = x
        }
    }

    return
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}