func findMaximumLength(nums []int) int {
    n := len(nums)
    s := make([]int, n + 1)
    last := make([]int, n + 1)
    f := make([]int, n + 1)
    q := make([]int, n + 1)

    front, rear := 0, 0 
    for i := 1; i <= n; i++ {
        s[i] = s[i - 1] + nums[i - 1]
        for front < rear && s[q[front + 1]] + last[q[front + 1]] <= s[i] + last[i] {
            front++
        }

        f[i] = f[q[front]] + 1
        last[i] = s[i] - s[q[front]]
        
        for rear >= front && s[q[rear]] + last[q[rear]] > s[i] + last[i] {
            rear--
        }
        rear++ 
        q[rear] = i 
    }

    return f[n]
}