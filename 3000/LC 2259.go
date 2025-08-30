func removeDigit(number string, digit byte) string {
    n := len(number)
    pre := -1

    for i := range number {
        if number[i] == digit {
            pre = i
            if i + 1 < n && number[i + 1] > number[i] {
                break
            }
        }
    }

    return number[:pre] + number[pre + 1:]
}