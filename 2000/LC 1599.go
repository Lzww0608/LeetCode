func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) (ans int) {
    ans = -1
    mx := 0
    sum := 0
    cur := 0
    for i, x := range customers {
        sum += x
        cur += min(4, sum) * boardingCost - runningCost
        sum -= min(4, sum)

        if mx < cur {
            ans = i + 1
            mx = cur
        }
    }

    cnt := len(customers)
    for sum > 0 && min(4, sum) *boardingCost - runningCost > 0 {
        cur += min(4, sum) * boardingCost - runningCost
        sum -= min(4, sum)

        if mx < cur {
            ans = cnt + 1
            mx = cur
        }
        cnt++
    }

    return 
}