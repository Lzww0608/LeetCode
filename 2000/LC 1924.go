const EPS = 1e-6

type point struct {
    x, y float64
}

func dis(p, q point) float64 {
    return (p.x - q.x) * (p.x - q.x) + (p.y - q.y) * (p.y - q.y)
}

func circumcenter(a, b, c point) point {
    a1, b1 := b.x - a.x, b.y - a.y
    a2, b2 := c.x - a.x, c.y - a.y 
    c1, c2 := a1 * a1 + b1 * b1, a2 * a2 + b2 * b2 
    d := 2 * (a1 * b2 - a2 * b1)
    return point{
        x: a.x + (c1 * b2 - c2 * b1) / d, 
        y: a.y + (a1 * c2 - a2 * c1) / d,
    }
}

func outerTrees(trees [][]int) []float64 {
    a := make([]point, len(trees))
    for i, tree := range trees {
        a[i] = point{float64(tree[0]), float64(tree[1])}
    }

    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(a), func(i, j int) {a[i], a[j] = a[j], a[i]})
    o, r := a[0], 0.0

    for i, p := range a {
        if dis(o, p) < r + EPS {
            continue
        }
        o, r = p, 0.0 
        for j, q := range a[:i] {
            if dis(o, q) < r + EPS {
                continue
            }
            o = point{(p.x + q.x) / 2, (p.y + q.y) / 2}
            r = dis(o, p)
            for _, x := range a[:j] {
                if dis(o, x) > r + EPS {
                    o = circumcenter(p, q, x)
                    r = dis(o, p)
                }
            }
        }
    }

    return []float64{o.x, o.y, math.Sqrt(r)}
}