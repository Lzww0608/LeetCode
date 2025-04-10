func minimumIndex(nums []int) int {
    mx, cnt := -1, 0
    for _, x := range nums {
        if x == mx {
            cnt++
        } else {
            cnt--
            if cnt < 0 {
                cnt = 1
                mx = x 
            }
        }
    }

    cnt = 0
    for _, x := range nums {
        if x == mx {
            cnt++
        }
    }

    n := len(nums)
    cur := 0
    for i := 0; i < n - 1; i++ {
        if nums[i] == mx {
            cur++
        }

        if cur * 2 > i + 1 && (cnt - cur) * 2 > n - i - 1 {
            return i 
        }
    }

    return -1
}