func maximumAND(nums []int, k int, m int) (ans int) {
    n := len(nums)
    a := make([]int, n)

    for b := 30; b >= 0; b-- {
        target := ans | (1 << b)
        for i, x := range nums {
            j := 30
            for ; j >= 0; j-- {
                if target & (1 << j) == 0 {
                    continue
                } else if x & (1 << j) == 0 {
                    break 
                } 
            }
            mask := (1 << (j + 1)) - 1
            a[i] = (target & mask) - (x & mask)
        }
        sort.Ints(a)
        sum := 0
        for i := range m {
            sum += a[i]
        }

        if sum <= k {
            ans = target
        }
    }

    return
}