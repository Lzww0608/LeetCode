func countGoodTriplets(arr []int, a int, b int, c int) (ans int) {
    n := len(arr)
    mx := slices.Max(arr)
    s := make([]int, mx + 2)
    for j, y := range arr {
        for k := j + 1; k < n; k++ {
            z := arr[k]
            if abs(y - z) > b {
                continue
            }
            l := max(y - a, z - c, 0)
            r := min(y + a, z + c, mx)
            ans += max(s[r+1] - s[l], 0)
        }

        for x := y + 1; x <= mx + 1; x++ {
            s[x]++
        }
    }

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}