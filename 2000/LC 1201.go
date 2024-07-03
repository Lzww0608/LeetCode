func getKthMagicNumber(k int) int {
    nums := make([]int, k)
    nums[0] = 1
    a, b, c := 0, 0, 0
    for i := 1; i < k; i++ {
        x, y, z := nums[a] * 3, nums[b] * 5, nums[c] * 7
        t := min(x, y, z)
        nums[i] = t
        if t == x {
            a++
        } func nthUglyNumber(n int, a int, b int, c int) int {
    l, r := 1, int(2*1e9)
    ab := a * b / gcd(a, b)
    bc := b * c / gcd(b, c)
    ac := a * c / gcd(a, c)
    abc := ab * c / gcd(ab, c)
    for l <= r {
        mid := l + ((r - l) >> 1)
        s := mid / a + mid / b + mid / c - mid / ab - mid / ac - mid / bc + mid / abc 
        if s >= n {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    return l
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}
        if t == y {
            b++
        }
        if t == z {
            c++
        }
    }

    return nums[k-1]
}