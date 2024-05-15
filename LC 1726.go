func tupleSameProduct(nums []int) (ans int) {
    m := map[int]int{}
    for i := range nums {
        for j := 0; j < i; j++ {
            m[nums[i] * nums[j]]++
        }
    }
    for _, v := range m {
        ans += (v - 1) * v * 4
    }

    return 
}