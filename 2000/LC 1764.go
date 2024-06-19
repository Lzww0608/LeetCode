func canChoose(groups [][]int, nums []int) bool {
    n, m := len(groups), len(nums)
    i, j := 0, 0
    for j < m && i < n {
        k := 0
        for j < m && k < len(groups[i]) && groups[i][k] == nums[j] {
            j++
            k++
        }
        if k == len(groups[i]) {
            i++
        } else {
            j = j - k + 1
        }

    }

    return i == n
}
