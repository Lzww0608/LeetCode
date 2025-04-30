var monthLength []int

func isLeapYear(year int) bool {
    return ((year % 400) == 0 || (year % 100 != 0 && year % 4 == 0))
}

func dateToInt(date string) int {
    var year, month, day int
    s := strings.Split(date, "-")
    year, _ = strconv.Atoi(s[0])
    month, _ = strconv.Atoi(s[1])
    day, _ = strconv.Atoi(s[2])

    ans := day - 1
    for month != 0 {
        month--
        ans += monthLength[month]
        if month == 2 && isLeapYear(year) {
            ans++
        }
    }

    ans += 365 * (year - 1971)
    ans += (year - 1) / 4 - 1971 / 4 
    ans -= (year - 1) / 100 - 1971 / 100
    ans += (year - 1) / 400 - 1971 / 400
    return ans
}

func daysBetweenDates(date1 string, date2 string) int {
    monthLength = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

    return abs(dateToInt(date1) - dateToInt(date2))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}