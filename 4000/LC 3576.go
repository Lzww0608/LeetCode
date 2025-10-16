func canMakeEqual(nums []int, k int) bool {
    cnt := 0
    n := len(nums)
    for _, x := range nums {
        if x == 1 {
            cnt++
        }
    }

    if cnt & 1 == 1 && (n - cnt) & 1 == 1 {
        return false
    }

    if cnt & 1 == 0 {
        sum, pre := 0, -1
        for i, x := range nums {
            if x == 1 {
                if pre == -1 {
                    pre = i 
                } else {
                    sum += i - pre 
                    pre = -1
                }
            }
        }
        if sum <= k {
            return true
        }
    }

    if (n - cnt) & 1 == 0 {
        sum, pre := 0, -1
        for i, x := range nums {
            if x == -1 {
                if pre == -1 {
                    pre = i 
                } else {
                    sum += i - pre 
                    pre = -1
                }
            }
        }
        if sum <= k {
            return true
        }
    }

    return false
}