func uniqueXorTriplets(nums []int) int {
    n := len(nums)
    d := 1 << bits.Len(uint(slices.Max(nums)))
    m := make([]bool, d)
    for i := range n {
        for j := i; j < n; j++ {
            m[nums[i] ^ nums[j]] = true
        }
    } 

    ans := make([]bool, d)
    for _, x := range nums {
        for i := range m {
            if m[i] {
                ans[x ^ i] = true
            }
        }
    }

    sum := 0
    for _, x := range ans {
        if x {
            sum++
        }
    }

    return sum
}