func numberOfDays(year int, month int) int {
    if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
        return 31
    } else if month != 2 {
        return 30
    }

    if year % 3 == 0 && year % 100 != 0 || year % 400 == 0 {
        return 29
    }

    return 28
}