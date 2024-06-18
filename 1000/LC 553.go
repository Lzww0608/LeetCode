func optimalDivision(nums []int) string {
    n := len(nums)
    buider := strings.Builder{}
    buider.WriteString(strconv.Itoa(nums[0]))
    if n == 1 {
        return buider.String()
    }
    if n == 2 {
        return buider.String() + "/" + strconv.Itoa(nums[1])
    }
    buider.WriteString("/(")
    for i := 1; i < n; i++ {
        buider.WriteString(strconv.Itoa(nums[i]))
        if i < n - 1 {
            buider.WriteString("/")
        }
    }
    buider.WriteString(")")

    return buider.String()
}