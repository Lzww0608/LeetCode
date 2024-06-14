func sumOddLengthSubarrays(arr []int) (ans int) {
    //贡献法
    n := len(arr)
    
    for i, x := range arr {
        leftEven, leftOdd := i / 2 + 1, (i + 1) / 2
        rightEven, rightOdd := (n - i - 1) / 2 + 1, (n - i) / 2
        ans += x * (leftEven * rightEven + leftOdd * rightOdd)
    }

    return 
}
