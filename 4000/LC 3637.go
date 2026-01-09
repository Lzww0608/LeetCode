func isTrionic(nums []int) bool {
    cnt := 0
    n := len(nums)
    var i int 
    for i = 1; i < n && cnt < 3; i++ {
        if nums[i] == nums[i - 1] {
            return false
        }
        if cnt & 1 == 0 {
            if nums[i] < nums[i - 1] {
                if i == 1 {
                    return false
                }
                cnt++
            } 
        } else {
            if nums[i] > nums[i - 1] {
                cnt++
            }
        }
    }

    return cnt == 2 && i == n
}