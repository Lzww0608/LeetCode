func maxNonDecreasingLength(a []int, b []int) (ans int) {
    n := len(a)
    f, g := 1, 1
    ans = 1
    for i := 1; i < n; i++ {
        x, y := f, g 
        f, g = 1, 1
        if a[i] >= a[i-1] {
            f = max(f, x + 1)
        }
        if a[i] >= b[i-1] {
            f = max(f, y + 1)
        }

        if b[i] >= a[i-1] {
            g = max(g, x + 1)
        }
        if b[i] >= b[i-1] {
            g = max(g, y + 1)
        }

        ans = max(ans, f, g)
    }
   
    return
}