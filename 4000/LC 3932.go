func countKthRoots(l int, r int, k int) (ans int) {
    if k == 1 {
        return r - l + 1
    }
    for i := 0; ; i++ {
        x := pow(i, k)
        if x > r {
            break
        } else if x >= l {
            ans++
        }
    }

    return
}

func pow(a, r int) int {
    if a <= 1 {
        return a
    }
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a
        }

        a = a * a 
        r >>= 1
    }

    return res
}