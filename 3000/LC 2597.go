func beautifulSubsets(nums []int, k int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    a := make([]int, 0, n)
    for i := 0; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            if nums[j] - nums[i] == k {
                a = append(a, (1 << i) | (1 << j))
            }
        }
    }

    for i := 1; i < (1 << n); i++ {
        f := true
        for _, x := range a {
            if i & x == x {
                f = false
                break
            }
        }
        if f {
            ans++
        }
    }

    return
}