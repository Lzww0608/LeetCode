func reformatDate(date string) string {
    s := strings.Split(date, " ")
    ans := s[2] + "-"

    ans += solve(s[1]) + "-"
    if len(s[0]) == 3 {
        ans += "0" + s[0][:1]
    } else {
        ans += s[0][:2]
    }

    return ans
}

func solve(s string) string {
    if s == "Jan" {
        return "01"
    } else if s == "Feb" {
        return "02"
    } else if s == "Mar" {
        return "03"
    } else if s == "Apr" {
        return "04"
    } else if s == "May" {
        return "05"
    } else if s == "Jun" {
        return "06"
    } else if s == "Jul" {
        return "07"
    } else if s == "Aug" {
        return "08"
    } else if s == "Sep" {
        return "09"
    } else if s == "Oct" {
        return "10"
    } else if s == "Nov" {
        return "11"
    } else {
        return "12"
    }
}