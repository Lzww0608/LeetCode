const MOD = int64(1e9 + 7)
func rectangleArea(rectangles [][]int) int {
    list := []int{}
    for _, info := range rectangles {
        list = append(list, info[0], info[2])
    }
    sort.Ints(list)
    ans := int64(0)
    n := len(list)
    for i := 1; i < n; i++ {
        a, b, length := list[i-1], list[i], list[i] - list[i-1]
        if length == 0 {
            continue
        }
        lines := [][]int{}
        for _, info := range rectangles {
            if info[0] <= a && b <= info[2] {
                lines = append(lines, []int{info[1], info[3]});
            }
        }
        sort.Slice(lines, func(i, j int) bool {
        if lines[i][0] == lines[j][0] {
                return lines[i][1] < lines[j][1]
            }
            return lines[i][0] < lines[j][0]
        })
        total, l, r := int64(0), -1, -1
        for _, info := range lines {
            if info[0] > r {
                total += int64(r - l)
                l, r = info[0], info[1]
            } else if info[1] > r {
                r = info[1]
            }
        }
        total += int64(r - l) 
        ans += int64(total) * int64(length)
        ans %= MOD
    }
    return int(ans)
}
