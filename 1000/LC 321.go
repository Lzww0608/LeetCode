func maxSequence(a []int, k int) (s []int) {
    for i, v := range a {
        for len(s) > 0 && len(s) + len(a) - i - 1 >= k && v > s[len(s)-1] {
            s = s[:len(s)-1]
        }

        if len(s) < k {
            s = append(s, v)
        }
    }

    return
}

func cmp(a, b []int) bool {
    for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] != b[i] {
            return a[i] < b[i]
        }
    }

    return len(a) < len(b)
}

func merge(a, b []int) []int {
    ans := make([]int, len(a) + len(b))
    for i := range ans {
        if cmp(a, b) {
            ans[i], b = b[0], b[1:]
        } else {
            ans[i], a = a[0], a[1:]
        }
    }

    return ans
}

func maxNumber(nums1 []int, nums2 []int, k int) (ans []int) {
    start := max(0, k - len(nums2))

    for i := start; i <= k && i <= len(nums1); i++ {
        a := maxSequence(nums1, i)
        b := maxSequence(nums2, k - i)
        tmp := merge(a, b)
        if cmp(ans, tmp) {
            ans = tmp
        }
    }

    return
}
