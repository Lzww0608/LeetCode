const N int = 256
func sampleStats(count []int) (ans []float64) {
    // minimum
    for i, x := range count {
        if x != 0 {
            ans = append(ans, float64(i))
            break
        }
    }
    // maximum
    for i := N - 1; i >= 0; i-- {
        if count[i] != 0 {
            ans = append(ans, float64(i))
            break
        }
    }
    // mean
    var sum float64
    var cnt int64
    for i, x := range count {
        cnt += int64(x)
        sum += float64(i * x)
    }
    ans = append(ans, float64(sum) / float64(cnt))
    // median
    a := -1
    var s int64 = 0
    //fmt.Println(cnt)
    for i, x := range count {
        if x == 0 {
            continue
        }
        if a != -1 {
            ans = append(ans, float64(a + i) / 2)
            break
        } else if s + int64(x) == cnt / 2 && cnt & 1 == 0 {
            a = i
        } else if s + int64(x) >= (cnt + 1) / 2 {
            ans = append(ans, float64(i))
            break
        } 
        s += int64(x)
    }
    // mode
    mx, t := 0, 0
    for i, x := range count {
        if x > mx {
            mx, t = x, i
        }
    }
    ans = append(ans, float64(t))

    return
}