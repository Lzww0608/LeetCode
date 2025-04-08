func minimumOperations(nums []int) (ans int) {
    m := make(map[int]int)
    for i, j := 0, 0; i < len(nums); i++ {
        m[nums[i]]++;
        for m[nums[i]] > 1 {
            for k := 0; k < 3 && j < len(nums); k++ {
                m[nums[j]]--
                j++
            }
            ans++
        }
    }

    return
}