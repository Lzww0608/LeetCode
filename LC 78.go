func subsets(nums []int) [][]int {
    ans := [][]int{}
    sub := []int{}
    n := len(nums)
    var solve func(int)
    solve = func(id int) {
        //tmp := make([]int, len(sub))
        //copy(tmp, sub)
        //ans = append(ans, tmp)
        ans = append(ans, append([]int{}, sub...))
        for i := id; i < n; i++ {
            sub = append(sub, nums[i])
            solve(i + 1)
            sub = sub[:len(sub)-1]
        }
    }
    solve(0)
    return ans
}