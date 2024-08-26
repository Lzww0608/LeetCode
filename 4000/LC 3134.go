func medianOfUniquenessArray(nums []int) int {
    n := len(nums)
    k := (n * (n + 1) / 2 + 1) / 2
    ans := 1 + sort.Search(n - 1, func(idx int) bool {
        idx++
        cnt, l := 0, 0
        fre := make(map[int]int)
        
        for r, in := range nums {
            fre[in]++
            for len(fre) > idx {
                out := nums[l]
                fre[out]--
                if fre[out] == 0 {
                    delete(fre, out)
                }
                l++
            }
            cnt += r - l + 1
            if cnt >= k {
                return true
            }
        }
        
        return false
    })

    return ans
}