func countElements(nums []int, k int) (ans int) {
    n := len(nums)
    if k == 0 {
        return n
    }
    sort.Ints(nums)

    for i, j := 0, 0; j + k <= n; i++ {
        for j < n && nums[j] == nums[i] {
            j++
        }

        if j >= n || j + k > n {
            break
        }
        ans++
    }

    return
}