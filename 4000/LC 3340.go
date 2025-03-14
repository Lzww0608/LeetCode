func isBalanced(num string) bool {
    sum := 0
    for i := range num {
        x := int(num[i] - '0')
        if i & 1 == 0 {
            sum += x
        } else {
            sum -= x
        }
    } 

    return sum == 0
}