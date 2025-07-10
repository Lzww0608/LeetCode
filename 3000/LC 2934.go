func minOperations(a []int, b []int) int {
    n := len(a)
    x := solve(a, b)
    if x == -1 {
        return -1
    }
    a[n-1], b[n-1] = b[n-1], a[n-1]
    y := solve(a, b) + 1

    return min(x, y)
}

func solve(a, b []int) (ans int) {
    n := len(a)

    for i := 0; i < n - 1; i++ {
        if a[i] > a[n-1] || b[i] > b[n-1] {
            if a[i] > b[n-1] || b[i] > a[n-1] {
                return -1
            }
            ans++
        }

        
    }

    return 
}