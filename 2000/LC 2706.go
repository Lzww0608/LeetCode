func buyChoco(prices []int, money int) int {
    sort.Ints(prices)
    if money >= prices[0] + prices[1] {
        return money - prices[0] - prices[1]
    }

    return money
}