const MOD int = int(1e9 + 7)
func firstDayBeenInAllRooms(nextVisit []int) int {
    n := len(nextVisit)
    //s := make([]int, n)
    j := nextVisit[0]
    for i := range nextVisit[:n-1] {
        nextVisit[i+1], j = (nextVisit[i] * 2 - nextVisit[j] + 2 + MOD) % MOD, nextVisit[i+1]
    }

    return nextVisit[n-1]
}