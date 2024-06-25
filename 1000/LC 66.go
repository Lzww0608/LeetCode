func plusOne(digits []int) []int {
    n := len(digits)
    i, add := n - 1, 1
    for i >= 0 {
        t := digits[i]
        if t + add == 10 {
            add = 1
            digits[i] = 0
        } else {
            digits[i] = t + 1
            return digits
        }
        i--
    } 

    digits = append([]int{1}, digits...)

    return digits
}
