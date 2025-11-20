func missingInteger(nums []int) int {
    m := make(map[int]bool)
    sum, d := nums[0], 1

    cur := 1
    for _, x := range nums {
        m[x] = true
    } 
    for i := 1; i < len(nums); i++ {
        x := nums[i]
        if nums[i] == nums[i - 1] + 1 {
            cur++
        } else {
            break
        }

        if cur > d || d == cur && sum < cur * (x + x - cur + 1) / 2 {
            sum = cur * (x + x - cur + 1) / 2
            d = cur
        } 
    }

    ans := sum
    for m[ans] {
        ans++
    }
    
    return ans
}