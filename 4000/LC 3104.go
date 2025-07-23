const N int = 26
func maxSubstringLength(s string) (ans int) {
    ans = -1
    n := len(s)
    left := make([]int, N)
    right := make([]int, N)
    for i := range N {
        left[i] = -1
    }

    for i := range s {
        x := int(s[i] - 'a')
        if left[x] == -1 {
            left[x] = i
        } 
        right[x] = i
    }

    intervals := [][2]int{}
    for i, l := range left {
        if l == -1 {
            continue
        }
        r := right[i]
        for j := l; j < r; j++ {
            x := int(s[j] - 'a')
            if left[x] < l {
                l = -1
                break
            }
            r = max(r, right[x])
        }

        if l != -1 && r - l + 1 != n {
            intervals = append(intervals, [2]int{l, r})
        }
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    for i, v := range intervals {
        l, r := v[0], v[1]
        ans = max(ans, r - l + 1)
        for j := i + 1; j < len(intervals); j++ {
            if intervals[j][0] == r + 1 {
                r = intervals[j][1]
                if r - l + 1 < n {
                    ans = max(ans, r - l + 1)
                }
            }
        }
    }

    return
}