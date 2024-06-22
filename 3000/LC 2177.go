func sumOfThree(num int64) []int64 {
    if num % 3 != 0 {
        return nil
    }
    
    return []int64{num / 3 - 1, num / 3, num / 3 + 1}
}
