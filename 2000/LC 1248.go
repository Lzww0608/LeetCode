func numberOfSubarrays(nums []int, k int) (ans int) {
    a := make([]int, 0, len(nums))
    a = append(a, -1)
    for i, x := range nums {
        if x & 1 == 1 {
            a = append(a, i)
        }
    }

    a = append(a, len(nums))
    n := len(a)
    for l, r := 1, k; r < n - 1; r++ {
        ans += (a[l] - a[l-1]) * (a[r+1] - a[r])
        l++
    }

    return
}