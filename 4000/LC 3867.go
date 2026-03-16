func gcdSum(nums []int) int64 {
    n := len(nums)
    mx := 0
    a := make([]int, n)
    for i, x := range nums {
        mx = max(mx, x)
        g := gcd(x, mx)
        a[i] = g
    }
    sort.Ints(a)
    
    ans := 0
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        ans += gcd(a[i], a[j])
    }

    return int64(ans)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}