const N int = 100
func singleDivisorTriplet(nums []int) int64 {
    cnt := [N + 1]int{}
    for _, x := range nums {
        cnt[x]++
    }

    ans := 0
    for i := 1; i <= N; i++ {
        if cnt[i] == 0 {
            continue
        }
        x := cnt[i]
        for j := i; j <= N; j++ {
            y := cnt[j]
            if y == 0 {
                continue
            }
            for k := j; k <= N; k++ {
                z := cnt[k]
                if z == 0 {
                    continue
                }
                sum := i + j + k 
                t := 0
                for _, v := range []int{i, j, k} {
                    if sum % v == 0 {
                        t++
                    }
                } 
                if t == 1 {
                    if i == j {
                        if j == k {
                            ans += x * (y - 1) * (z - 2)
                        } else {
                            ans += x * (x - 1) * z * 3
                        }
                    } else if j == k {
                        ans += x * y * (z - 1) * 3
                    } else {
                        ans += x * y * z * 6
                    }
                }
            }
        }
    }

    return int64(ans)
}