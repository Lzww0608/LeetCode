func maxDistance(side int, points [][]int, k int) int {
    n := len(points)
    a := make([]int, 0, n)
    for _, v := range points {
        x, y := v[0], v[1]
        if x == 0 {
            a = append(a, y)
        } else if y == side {
            a = append(a, side + x)
        } else if x == side {
            a = append(a, side * 3 - y)
        } else {
            a = append(a, side * 4 - x)
        }
    }
    sort.Ints(a)

    f := make([]int, n + 1)
    end := make([]int, n)
    
    check := func(mid int) bool {
        j := n
        for i := n - 1; i >= 0; i-- {
            for a[j - 1] >= a[i] + mid {
                j--
            }

            f[i] = f[j] + 1
            if f[i] == 1 {
                end[i] = i
            } else {
                end[i] = end[j]
            }

            for f[i] == k && a[end[i]] - a[i] <= side * 4 - mid {
                return true
            }
        }

        return false
    }

    l, r := 1, side * 4 / k + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid
        } else {
            r = mid 
        }
    }

    return l
}