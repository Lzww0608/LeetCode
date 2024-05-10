func sequenceReconstruction(nums []int, sequences [][]int) bool {
    n := len(nums)
    m := make([]int, n + 1)
    for i := 0; i < n - 1; i++ {
        m[nums[i]] = nums[i+1]
    }

    cnt := n - 1
    for _, a := range sequences {
        for i := 0; i < len(a) - 1; i++ {
            if m[a[i]] == a[i+1] {
                m[a[i]] = -1
                cnt--
            } else if m[a[i]] == 0{
                return false
            }
        }

    }

    return cnt == 0
}