func numGoodSubarrays(nums []int, k int) int64 {
    ans := 0
    sum := 0

    m := make(map[int]int)
    m[0] = 1
    for _, x := range nums {
        sum += x 
        ans += m[sum % k]
        m[sum % k]++
    }

    n := len(nums)
    a := []int{}
    for i := 0; i < n; i++ {
        a = a[:0]
        sum = 0
        j := i 

        for j < n && nums[j] == nums[i] {
            sum += nums[j]
            if sum % k == 0 {
                a = append(a, j - i + 1)
            }
            j++
        }
        d := j - i
        i = j - 1

        for _, x := range a {
            ans -= d - x
        }
    }

    return int64(ans)
}