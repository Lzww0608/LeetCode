func minimumNumbers(num int, k int) int {
    if num == 0 {
        return 0
    }
    for i := 1; i <= 10; i++ {
        if (num - i * k) >= 0 && (num - i * k) % 10 == 0 {
            return i
        }
    }

    return -1
}