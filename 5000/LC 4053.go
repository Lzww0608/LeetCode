var a, b []int

func init() {
    for i := 0; i < 100_001; i++ {
        s := []byte(strconv.Itoa(i))
        t := make([]byte, len(s))
        copy(t, s)
        n := len(s)
        for i := n - 1; i >= 0; i-- {
            s = append(s, s[i])
        }
        for i := n - 2; i >= 0; i-- {
            t = append(t, t[i])
        }
        
        x, _ := strconv.Atoi(string(s))
        y, _ := strconv.Atoi(string(t))
        if x & 1 == 0 {
            a = append(a, x)
        } else {
            b = append(b, x)
        }

        if y & 1 == 0 {
            a = append(a, y)
        } else {
            b = append(b, y)
        }
    }
    sort.Ints(a)
    sort.Ints(b)
}

func minOperations(nums []int) int64 {
    ans := 0
    for _, x := range nums {
        t := 0
        if x & 1 == 0 {
            p := sort.SearchInts(a, x)
            if p == 0 {
                t = a[p] - x
            } else {
                t = min(a[p] - x, x - a[p - 1])
            }
        } else {
            p := sort.SearchInts(b, x)
            if p == 0 {
                t = b[p] - x
            } else {
                t = min(b[p] - x, x - b[p - 1])
            }
        }
        ans += t / 2
        //fmt.Println(ans)
    }

    return int64(ans)
}