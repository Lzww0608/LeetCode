func lastRemaining(n int) int {
    head, step := 1, 1
    left := true 

    for n > 1 {
        if left {
            head += step
            left = false
        } else {
            if n & 1 == 1 {
                head += step
            }
            left = true
        }

        n /= 2
        step <<= 1
    }

    return head
}