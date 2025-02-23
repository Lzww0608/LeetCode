func minCost(nums []int, k int) int {
    n := len(nums)
    f := make([]int, n + 1)
    cnt := make(map[int]int)
    sum := 0
    for i, x := range nums {
        if cnt[x]++; cnt[x] > 1 {
            if cnt[x] == 2 {
                sum += 2
            } else {
                sum++
            }
        } 
        f[i+1] = min(sum + k, f[i] + k)

        cur := make(map[int]int)
        t := 0 
        cur[x]++
        for j := i - 1; j >= 0; j-- {
            if cur[nums[j]]++; cur[nums[j]] > 1 {
                if cur[nums[j]] == 2 {
                    t += 2
                } else {
                    t++
                }
            }
            f[i+1] = min(f[i+1], f[j] + k + t)
        }
    }
    //fmt.Println(f)
    return f[n]
}