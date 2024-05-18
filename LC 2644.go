func maxDivScore(nums []int, divisors []int) (ans int) {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    cnt := -1
    for _, x := range divisors {
        t := 0
        for _, a := range nums {
            if a < x {
                break
            }
            if a % x == 0 {
                t++
            }
            
        }
        if t > cnt || (cnt == t && ans > x) {
            cnt, ans = t, x
        }
    } 
    return 
}