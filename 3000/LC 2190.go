func mostFrequent(nums []int, key int) (ans int) {
    m := make(map[int]int)

    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == key {
            m[nums[i + 1]]++
        }
    }

    mx := 0
    for k, v := range m {
        if v > mx {
            ans = k 
            mx = v
        }
    }

    return
}