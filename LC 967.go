func numsSameConsecDiff(n int, k int) []int {
    q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    for n > 1 {
        n--
        for i := len(q); i > 0; i-- {
            t := q[0]
            q = q[1:]
            num := t % 10
            if num + k < 10 {
                q = append(q, t * 10 + num + k)
            }
            if k != 0 && num - k >= 0 {
                q = append(q, t * 10 + num - k)
            }
        }
    }

    return q
}