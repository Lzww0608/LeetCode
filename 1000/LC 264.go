func nthUglyNumber(n int) int {
    ugly := make([]int, n)
    ugly[0] = 1
    a, b, c := 0, 0, 0
    for i := 1; i < n; i++ {
        x, y, z := ugly[a] * 2, ugly[b] * 3, ugly[c] * 5
        t := min(x, y, z)
        ugly[i] = t
        if t == x {
            a++
        } 
        if t == y {
            b++
        }
        if t == z {
            c++
        }
    }

    return ugly[n-1]
}