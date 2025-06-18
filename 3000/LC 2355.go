func maximumBooks(books []int) int64 {
    var ans int 
    var st [][2]int
    st = append(st, [2]int{-1, 0})

    for i, x := range books {
        for len(st) > 1 && books[st[len(st)-1][0]] - st[len(st)-1][0] > x - i {
            st = st[:len(st)-1]
        }

        d := min(x, i - st[len(st)-1][0])
        s := (x * 2 - d + 1) * d / 2 + st[len(st)-1][1]
        ans = max(ans, s)
        st = append(st, [2]int{i, s})
    }

    return int64(ans)
}