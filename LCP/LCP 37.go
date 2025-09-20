const EPS = 1e-8
const T int = 60
func minRecSize(lines [][]int) float64 {
    n := len(lines)

    binaryFind := func(c1, c2 bool) float64 {
        l, r := float64(-1e9), float64(1e9)

        for t := T; t > 0; t-- {
            mid := l + ((r - l) / 2.0)

            f := false
            for i := 1; i < n; i++ {
                a := float64(lines[i][0]) * mid + float64(lines[i][1])
                b := float64(lines[i - 1][0]) * mid + float64(lines[i - 1][1])
                if c1 {
                    a = (mid - float64(lines[i][1])) / float64(lines[i][0])
                b = (mid - float64(lines[i - 1][1])) / float64(lines[i - 1][0])
                }
                if b - a > EPS {
                    f = true
                    break
                }
            }

            if f {
                if c2 {
                    r = mid
                } else {
                    l = mid
                }
            } else {
                if c2 {
                    l = mid
                } else {
                    r = mid
                }
            }
        }

        return r
    }

    sort.Slice(lines, func(i, j int) bool {
        if lines[i][0] == lines[j][0] {
            return lines[i][1] < lines[j][1]
        }
        return lines[i][0] < lines[j][0]
    })
    r := binaryFind(false, false)

    sort.Slice(lines, func(i, j int) bool {
        if lines[i][0] == lines[j][0] {
            return lines[i][1] < lines[j][1]
        }
        return lines[i][0] > lines[j][0]
    })
    l := binaryFind(false, true)

    sort.Slice(lines, func(i, j int) bool {
        if lines[i][0] == lines[j][0] {
            return lines[i][1] > lines[j][1]
        }
        return lines[i][0] > lines[j][0]
    })
    u := binaryFind(true, false)

    sort.Slice(lines, func(i, j int) bool {
        if lines[i][0] == lines[j][0] {
            return lines[i][1] > lines[j][1]
        }
        return lines[i][0] < lines[j][0]
    })
    d := binaryFind(true, true)

    if r < l + EPS || u < d + EPS {
        return 0.0
    }

    return (r - l) * (u - d)
}