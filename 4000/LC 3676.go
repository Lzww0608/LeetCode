func bowlSubarrays(nums []int) int64 {
    ans := 0
    st := []int{}

    for _, x := range nums {
        for len(st) > 0 && st[len(st) - 1] < x {
            st = st[:len(st) - 1]
            if len(st) > 0 {
                ans++
            }
        }

        st = append(st, x)
    }

    return int64(ans)
}