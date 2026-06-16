func checkGoodInteger(n int) bool {
    digitSum, squareSum := 0, 0
    for n > 0 {
        x := n % 10 
        n /= 10
        digitSum += x 
        squareSum += x * x
    }

    return squareSum - digitSum >= 50
}