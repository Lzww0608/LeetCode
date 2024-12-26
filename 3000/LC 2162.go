
import "strconv"
func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
    targetSeconds = targetSeconds / 60 * 100 + targetSeconds % 60  
    anotherTargetSeconds := targetSeconds
    if targetSeconds % 100 < 40 && targetSeconds / 100 != 0 {
        anotherTargetSeconds = targetSeconds + 60 - 100
    }
    //fmt.Println(targetSeconds, anotherTargetSeconds)

    s := strconv.Itoa(targetSeconds)
    t := strconv.Itoa(anotherTargetSeconds)
    cur, cost := startAt, 0
    for i := range s {
        x := int(s[i] - '0')
        if cur != x {
            cost += moveCost
            cur = x 
        } 
        cost += pushCost
    }

    ans := cost
    cur, cost = startAt, 0 
    for i := range t {
        x := int(t[i] - '0')
        if cur != x {
            cost += moveCost
            cur = x 
        } 
        cost += pushCost
    }
    if len(s) > 4 {
        return cost
    }
    return min(ans, cost)
}