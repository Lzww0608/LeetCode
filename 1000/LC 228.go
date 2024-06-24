func summaryRanges(nums []int) []string {
    ans := []string{}
    n := len(nums)

    l, r := 0, 1
    builder := strings.Builder{}
    for r <= n {
       if r == n || nums[r] != nums[r - 1] + 1 {
            builder.Reset()
            builder.WriteString(strconv.Itoa(nums[l]))
            if l + 1 < r {
                builder.WriteString("->")
                builder.WriteString(strconv.Itoa(nums[r-1]))
            }
            ans = append(ans, builder.String())
            l = r
        } 
        r++
    }



    return ans
}
