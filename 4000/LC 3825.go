func longestSubsequence(nums []int) (ans int) {
    a := make([]int, 0, len(nums))
    for i := 0; i <= 30; i++ {
        mask := 1 << i 
        a = a[:0]
        for _, x := range nums {
            if x & mask != 0 {
                a = append(a, x)
            }
        }
        
        ans = max(ans, calc(a))
    }

    return
}

func calc(a []int) int {
    st := []int{}

    for _, x := range a {
        if len(st) > 0 && x <= st[len(st) - 1] {
            p := sort.SearchInts(st, x)
            st[p] = x
        } else {
            st = append(st, x)
        }
    }

    return len(st)
}