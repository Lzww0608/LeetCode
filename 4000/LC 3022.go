const N int = 29
func minOrAfterOperations(nums []int, k int) (ans int) {
    mask := 0
    for b := N; b >= 0; b-- {
        mask |= 1 << b 
        cnt, and := 0, -1
        for _, x := range nums {
            and &= x & mask 
            if and != 0 {
                cnt++
            } else {
                and = -1
            }
        }

        if cnt > k {
            ans |= 1 << b 
            mask ^= 1 << b 
        }
    }

    return ans
}