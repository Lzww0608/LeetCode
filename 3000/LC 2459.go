func sortArray(nums []int) int {
    n := len(nums)

    solve := func(pos int) (cnt int) {
        vis := make([]bool, n)
        for i := 0; i < n; i++ {
            if nums[i] != i && !vis[i] {
                cnt++

                x := i
                for !vis[x] {
                    cnt++
                    vis[x] = true 
                    x = nums[x]
                }
            }
        }
        if nums[pos] != pos {
            cnt -= 2
        }
        return
    }
    ans := solve(0)
    for i := 0; i < n; i++ {
        nums[i] = (nums[i] + n - 1) % n;
    }
    return min(ans, solve(n - 1))
}