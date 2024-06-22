func containsDuplicate(nums []int) bool {
    m := map[int]int{}

    for _, x := range nums {
        m[x]++
        if m[x] >= 2 {
            return true
        }
    }


    return false
}
