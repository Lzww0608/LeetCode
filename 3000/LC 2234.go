
func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
    ans, n := 0, len(flowers)
    sort.Slice(flowers, func(i, j int) bool {
        return flowers[i] > flowers[j]
    })

    cpt := 0
    if flowers[n-1] >= target {
        return int64(n) * int64(full)
    }
    pos, sum := 0, 0
    for i, x := range flowers {
        if x >= target {
            cpt++
            pos = i + 1
        } else {
            sum += x 
        }
    }

    p, j := target - 1, pos 
    for i := pos - 1; i < n; i, cpt = i + 1, cpt + 1 {
        if i >= pos {
            newFlowers -= int64(target - flowers[i])
        }
        
        if newFlowers < 0 {
            break
        }
        
        if i < n - 1 {
            for j <= i {
                sum -= flowers[j]
                j++
            }

            for int64(p * (n - j) - sum) > newFlowers {
                p--
                for flowers[j] > p {
                    sum -= flowers[j]
                    j++
                }
            }
            ans = max(ans, cpt * full + p * partial)
        } else {
            ans = max(ans, cpt * full)
        }
        
    }

    return int64(ans)
}