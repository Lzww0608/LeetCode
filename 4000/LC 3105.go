func longestMonotonicSubarray(nums []int) int {
    des, inc := 1, 1
    f := true
    cnt := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            if f {
                cnt++
                inc = max(inc, cnt)
            } else {
                f = true
                cnt = 2
                des = max(des, cnt)
            }
        } else if nums[i] < nums[i-1] {
            if !f {
                cnt++
                des = max(des, cnt)
            } else {
                f = false
                cnt = 2
                inc = max(inc, cnt)
            }
        } else {
            cnt = 1
        }
    }

    return max(des, inc)
}