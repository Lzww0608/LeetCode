func maximumMEX(nums []int) (ans []int) {
    mx := slices.Max(nums)
    pos := make([][]int, mx + 2)
    for i, x := range nums {
        pos[x] = append(pos[x], i)
    }

    last, pre := mx + 1, -1 
    for {
        i := 0
        tmp1, tmp2 := last, pre
        for i < tmp1 {
            for len(pos[i]) > 0 && pos[i][0] < tmp2 {
                pos[i] = pos[i][1:]
            }
            if len(pos[i]) == 0 {
                last = min(last, i)
                break
            }
            pre = max(pre, pos[i][0])
            pos[i] = pos[i][1:]
            if len(pos[i]) == 0 {
                last = min(last, i)
            }
            i++
        }
        if i == 0 {
            for _, v := range pos {
                if len(v) == 0 {
                    continue
                }
                p := sort.SearchInts(v, pre)
                for range len(v) - p {
                    ans = append(ans, 0)
                }
            }
            break
        }

        ans = append(ans, i)
    }

    return 
}