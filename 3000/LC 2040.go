func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
    n, m := len(nums1), len(nums2)
    pa, na := make([]int, 0, n), make([]int, 0, n)
    pb, nb := make([]int, 0, m), make([]int, 0, m)

    for _, x := range nums1 {
        if x < 0 {
            na = append(na, -x)
        } else {
            pa = append(pa, x)
        }
    }
    for _, x := range nums2 {
        if x < 0 {
            nb = append(nb, -x)
        } else {
            pb = append(pb, x)
        }
    }

    sort.Ints(pa)
    sort.Ints(pb)
    sort.Ints(na)
    sort.Ints(nb)

    check := func(mid int) bool {
        cnt := 0
        if mid >= 0 {
            cnt += len(pa) * len(nb) + len(pb) * len(na)
            if cnt >= int(k) {
                return true
            }
            for _, x := range pa {
                if x == 0 {
                    cnt += len(pb)
                } else {
                    cnt += sort.SearchInts(pb, mid / x + 1)
                }
            }
            for _, x := range na {
                cnt += sort.SearchInts(nb, mid / x + 1)
            }
        } else {
            mid = -mid
            for _, x := range pa {
                if x != 0 {
                    cnt += len(nb) - sort.SearchInts(nb, (mid + x - 1) / x)
                }
                
            }
            for _, x := range na {
                cnt += len(pb) - sort.SearchInts(pb, (mid + x - 1) / x)
            }
        }

        return cnt >= int(k)
    }

    l, r := int(-2e10), int(2e10)
    for l + 1 < r {
        mid := div(l + r, 2)
        if check(mid) {
            r = mid
        } else {
            l = mid 
        }
    }

    return int64(r)
}

func div(a, b int) int {
    ans := a / b 
    if a ^ b < 0 && a % b != 0 {
        ans--
    }

    return ans
}