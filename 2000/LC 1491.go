func average(salary []int) float64 {
    n := len(salary)
    mn, mx := salary[0], salary[0]
    ans := float64(0)

    for _, x := range salary {
        ans += float64(x)
        mx = max(mx, x)
        mn = min(mn, x)
    }

    ans -= float64(mx + mn)
    return ans / float64(n - 2)
}
