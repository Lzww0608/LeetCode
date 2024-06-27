func majorityElement(nums []int) int {
    cnt, ans := 0, -1
    for _, x := range nums {
        if x != ans {
            cnt--
            if cnt < 0 {
                ans, cnt = x, 1
            }
        } else {
            cnt++
        }
    }

    if cnt > 1 {
        return ans
    } else if cnt == 1 && len(nums) & 1 == 1 {
        cnt = 0
        for _, x := range nums {
            if x == ans {
                cnt++
            }
        }
        if cnt > len(nums) / 2 {
            return ans
        } else {
            return -1
        }
    }

    return -1
}