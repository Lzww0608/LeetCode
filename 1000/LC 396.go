func maxRotateFunction(nums []int) int {
    n := len(nums)
    sum := 0
    tot := 0
    for i, x := range nums {
        sum += i * x
        tot += x 
    }

    ans := sum
    cur := sum  
    for _, x := range nums[:n-1] {
        cur = cur + n * x - tot 
        ans = max(ans, cur)
    }

    return ans
}