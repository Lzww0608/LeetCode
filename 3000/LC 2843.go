func countSymmetricIntegers(low int, high int) (ans int) {
    for i := low; i <= high; i++ {
        if solve(i) {
            ans++
        }
    }

    return
}

func solve(x int) bool {
    a := []int{}
    for x > 0 {
        a = append(a, x % 10)
        x /= 10
    }
    if len(a) & 1 == 1 {
        return false
    }

    sum, n := 0, len(a)
    for i := 0; i < n; i++ {
        if i < n / 2 {
            sum += a[i]
        } else {
            sum -= a[i]
        }
    }

    return sum == 0
}