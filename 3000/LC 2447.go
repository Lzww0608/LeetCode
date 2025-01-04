func subarrayGCD(nums []int, k int) (ans int) {
    n := len(nums)
    for i, x := range nums {
        cur := x
        if x == k {
            ans++
        }

        for j := i + 1; j < n; j++ {
            cur = gcd(cur, nums[j])
            if cur == k {
                ans++
            } else if cur % k != 0 {
                break
            }
        }
    }

    return
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}