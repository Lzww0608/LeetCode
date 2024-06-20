func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
        return a + b > b + a
    })
    
    builder := strings.Builder{}
    for _, x := range nums {
        builder.WriteString(strconv.Itoa(x))
    }

    s := builder.String()
    if len(s) > 1 && s[0] == '0' {
        return "0"
    }

    return s
}
