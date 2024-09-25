func minimumOperations(nums []int, start int, goal int) int {
    m := map[int]int{}

    q := [][]int{{start, 0}}
    for len(q) > 0 {
        cur, k := q[0][0], q[0][1]
        q = q[1:]
        for _, x := range nums {
            if cur + x == goal || cur - x == goal || cur ^ x == goal {
                return k + 1
            }
            if _, ok := m[cur + x]; !ok && cur + x >= 0 && cur + x <= 1000 {
                m[cur + x] = k
                q = append(q, []int{cur + x, k + 1})
            }
            if _, ok := m[cur - x]; !ok && cur - x >= 0 && cur - x <= 1000 {
                m[cur + x] = k
                q = append(q, []int{cur - x, k + 1})
            }
            if _, ok := m[cur ^ x]; !ok && cur ^ x >= 0 && cur ^ x <= 1000 {
                m[cur ^ x] = k
                q = append(q, []int{cur ^ x, k + 1})
            }
        }
        k++
    }

    return -1
}