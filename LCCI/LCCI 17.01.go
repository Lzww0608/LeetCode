func add(a int, b int) int {
    for b != 0 {
        a, b = a ^ b, (a & b) << 1
    }

    return a
}