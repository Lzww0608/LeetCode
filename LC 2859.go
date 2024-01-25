func sumIndicesWithKSetBits(nums []int, k int) int {
    countOnes := func(x int) int {
        cnt := 0
        for x != 0 {
            x = x & (x - 1)
            cnt++
            if cnt > k {
                return cnt
            }
        }
        return cnt
    }
    ans := 0
    for i, x := range nums {
        if countOnes(i) == k {
            ans += x
        }
    }
    return ans
}