func findRadius(houses []int, heaters []int) (ans int) {
    sort.Ints(houses)
    sort.Ints(heaters)
    j := 0

    for _, x := range houses {
        for j + 1 < len(heaters) && abs(heaters[j + 1] - x) <= abs(heaters[j] - x) {
            j++
        }

        ans = max(ans, abs(heaters[j] - x))
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}