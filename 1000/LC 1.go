func twoSum(nums []int, target int) []int {
    m := map[int]int{}

    for i, x := range nums {
        if idx, ok := m[target - x]; ok {
            return []int{idx, i}
        } else {
            m[x] = i
        }
    }

    return []int{-1, -1}
}