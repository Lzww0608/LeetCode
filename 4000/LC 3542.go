func minOperations(nums []int) (ans int) {
    st := []int{}
    for _, x := range nums {
        for len(st) > 0 && st[len(st) - 1] > x {
            ans++
            st = st[:len(st) - 1]
        }

        if x != 0 && (len(st) == 0 || st[len(st) - 1] != x) {
            st = append(st, x)
        }
    }

    ans += len(st)

    return
}