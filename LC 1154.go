func dayOfYear(date string) int {
    days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    s := strings.Split(date, "-")
    month := int(s[1][0] - '0') * 10 + int(s[1][1] - '0')
    day := int(s[2][0] - '0') * 10 + int(s[2][1] - '0')
    ans := day
    for i := 0; i < month; i++ {
        ans += days[i]
    }

    year := int(s[0][0] - '0') * 1000 + int(s[0][1] - '0') * 100 + int(s[0][2] - '0') * 10 + int(s[0][3] - '0')
    if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
        if month > 2 {
            ans++
        }
    }

    return ans
}