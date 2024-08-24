func countLatticePoints(circles [][]int) (ans int) {
    sort.Slice(circles, func(i, j int) bool {
        return circles[i][2] < circles[j][2]
    })

    for x := 0; x <= 200; x++ {
        for y := 0; y <= 200; y++ {
            for _, v := range circles {
                if (v[0] - x) * (v[0] - x) + (v[1] - y) * (v[1] - y) <= v[2] * v[2] {
                    ans++
                    break
                }
            }
        }
    }

    return 
}