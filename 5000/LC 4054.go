func shadowPairs(nums []int) int64 {
    st := [][]int{{0, 0}}
    cnt := 0
    ans := 0
    for _, x := range nums {
        for x < st[len(st) - 1][0] {
            cnt -= st[len(st) - 1][1]
            st = st[:len(st) - 1]
        }
        ans += cnt

        if st[len(st) - 1][0] == x {
            ans -= st[len(st) - 1][1]
            st[len(st) - 1][1] += 1
        } else {
            st = append(st, []int{x, 1})
        }

        cnt++
        fmt.Println(ans)
    }

    return int64(ans)
}