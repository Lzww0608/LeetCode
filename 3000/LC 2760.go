func longestAlternatingSubarray(nums []int, threshold int) (ans int) {
    cnt := 0
    if nums[0] <= threshold && nums[0] & 1 == 0 {
        cnt = 1
    }
    ans = cnt
    for i := 1; i < len(nums); i++ {
        if nums[i] > threshold {
            cnt = 0
            continue
        }

        if (nums[i] & 1) == (nums[i-1] & 1) {
            if nums[i] & 1 == 0 {
                cnt = 1
            } else {
                cnt = 0
            }
        } else if cnt > 0 || nums[i] & 1 == 0 {
            cnt++
        }

        ans = max(ans, cnt)
    }

    return
}