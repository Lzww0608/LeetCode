const N int = 26
func checkIfPangram(sentence string) bool {
    mask := 0
    for i := range sentence {
        mask |= 1 << int(sentence[i] - 'a')
    }

    return mask + 1 == 1 << N
}