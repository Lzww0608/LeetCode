const MOD int = 1_000_000_007
func maxArea(h int, w int, a []int, b []int) int {
    a = append(a, 0)
    a = append(a, h)
    b = append(b, 0)
    b = append(b, w)
    sort.Ints(a)
    sort.Ints(b)
    x, y := 0, 0 
    for i := 1; i < len(a); i++ {
        x = max(x, a[i] - a[i-1])
    }
    for i := 1; i < len(b); i++ {
        y = max(y, b[i] - b[i-1])
    }

    return x * y % MOD
}