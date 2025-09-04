
func makeArrayPositive(nums []int) (ans int) {
    mx := 0

    j := 0
    s := []int{0}
    for i, x := range nums {
        s = append(s, s[len(s) - 1] + x)
        i += 1
        if i - j > 2 && s[len(s) - 1] <= mx {
            ans++
            j = i 
            s[len(s) - 1] = 0
            mx = 0
        } else if i - j >= 2 {
            mx = max(mx, s[i - 2])
        }
    }

    return 
}