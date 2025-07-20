func makeSubKSumEqual(arr []int, k int) int64 {
    ans := 0
    n := len(arr)
    k = gcd(k, n)
    for i := 0; i < k; i++ {
        a := make([]int, 0, n)
        for j := i; j < n; j += k {
            a = append(a, arr[j])
        }
        ans += solve(a)
    }

    return int64(ans)
}

func solve(a []int) (ans int) {
    n := len(a)
    sort.Ints(a)
    for _, x := range a {
        ans += abs(x - a[n/2])
    }

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b 
    }

    return a
}