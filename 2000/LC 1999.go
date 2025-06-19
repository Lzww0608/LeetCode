func findInteger(k int, a int, b int) int {
    a, b = min(a, b), max(a, b)

    q := []int{a, b}
    for len(q) > 0 {
        cur := q[0]
        if cur > math.MaxInt32 {
            break
        }
        q = q[1:]
        if cur > k && cur % k == 0 {
            return cur
        }

        x := cur * 10 + a 
        y := cur * 10 + b 
        x, y = min(x, y), max(x, y)
        if x != cur {
            q = append(q, x)
        }
        if y != cur {
            q = append(q, y)
        }
    }

    return -1
}