func bestCoordinate(towers [][]int, radius int) []int {

    mx := -1
    x, y := 0, 0
    for i := 0; i <= 50; i++ {
        for j := 0; j <= 50; j++ {
            sum := 0
            for _, v := range towers {
                a, b, c := v[0], v[1], v[2]
                d := (a - i) * (a - i) + (b - j) * (b - j)
                if d > radius * radius {
                    continue
                }
                sum += int(float64(c) / (1 + math.Sqrt(float64(d))))
            }
            if sum > mx {
                x, y = i, j
                mx = sum
            } 
        }
    }

    return []int{x, y}
}
