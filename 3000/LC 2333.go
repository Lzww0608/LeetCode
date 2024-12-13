func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
    k := k1 + k2
    n := len(nums1)
    dif := make([]int, n)
    sum := 0
    for i := range nums1 {
        dif[i] = abs(nums1[i] - nums2[i])
        sum += dif[i]
    }

    if sum <= k {
        return 0
    }

    check := func(target int) bool {
        sum := 0
        for _, x := range dif {
            if x > target {
                sum += x - target
            }
        }
        return sum <= k 
    }

    l, r := 0, slices.Max(dif) + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid + 1 
        }
    }
    fmt.Println(dif, r)
    ans, cnt := 0, 0
    for _, x := range dif {
        if x < r {
            ans += x * x
        } else {
            k -= x - r
            ans += r * r
            cnt++
        }
    }

    ans += k * ((r - 1) * (r - 1) - r * r) 

    return int64(ans)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}