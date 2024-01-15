func circularArrayLoop(nums []int) bool {
    n := len(nums)
    if n < 2 {
        return false
    }

    next := func(i int) int {
        return ((i + nums[i])%n + n) % n // 确保结果在 0 到 n-1 的范围内
    }

    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            continue
        }

        // 使用快慢指针检测循环
        slow, fast := i, next(i)
        for nums[slow] * nums[fast] > 0 && nums[slow] * nums[next(fast)] > 0 {
            if slow == fast {
                // 检查循环是否长度大于1
                if slow == next(slow) {
                    break
                }
                return true
            }
            slow = next(slow)
            fast = next(next(fast))
        }

        // 标记已经访问过的元素
        slow = i
        for nums[slow] * nums[next(slow)] > 0 {
            nums[slow], slow = 0, next(slow)
        }
    }
    return false
}
