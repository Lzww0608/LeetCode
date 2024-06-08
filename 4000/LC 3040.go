func maxOperations(nums []int) int {

    var solve func(int, int, int) int
    solve = func(l, r, target int) (ans int) {
        if l >= r {
            return 0
        }
        if target == nums[l] + nums[r] {
            ans = max(ans, 1 + solve(l + 1, r - 1, target))
        }
        t := 0
        for l < r && target == nums[l] + nums[l+1] {
            t++
            l += 2
        }
        for l < r && target == nums[r] + nums[r-1] {
            t++
            r -= 2
        }
        if t > 0 {
            ans = max(ans, t + solve(l, r, target))
        }
        return
    }

    n := len(nums)
    a := solve(2, n - 1, nums[0] + nums[1])
    b := solve(1, n - 2, nums[0] + nums[n-1])
    c := solve(0, n - 3, nums[n-1] + nums[n-2])
    return max(a, b, c) + 1
}