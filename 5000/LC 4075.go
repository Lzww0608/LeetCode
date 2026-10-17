func resilientSubarray(nums []int, k int) (ans int) {
    sum := 0
    ans = 1
    p := make(map[int]int)
    for i, x := range nums {
        d := x % k 
        if i > 0 && d != nums[i - 1] % k {
            sum = 0
            clear(p)
        }
        sum += x 
        if j, ok := p[sum % k]; ok {
            ans = max(ans, i - j + 1)
        } else {
            p[sum % k] = i
        }
        //fmt.Println(sum)
    }

    return
}