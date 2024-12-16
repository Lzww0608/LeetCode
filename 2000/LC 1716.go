func totalMoney(n int) int {
    weeks := n / 7
    firstWeekMoney := (1 + 7) * 7 / 2
    lastWeekMoney := firstWeekMoney + 7 * (weeks - 1)
    weekMoney := (firstWeekMoney + lastWeekMoney) * weeks / 2

    days := n % 7 
    firstDayMoney := weeks + 1
    lastDayMoney := firstDayMoney + days - 1
    dayMoney := (firstDayMoney + lastDayMoney) * days / 2
    return weekMoney + dayMoney
}