const N int = 32
func onceTwice(nums []int) []int {
    cnt := [N]int{}
    for _, x := range nums {
        for i := 0; i < N; i++ {
            if x & (1 << i) != 0 {
                cnt[i]++
            }
        }
    }

    var a, b int32
    for i, x := range cnt {
        if x % 3 == 1 {
            a |= 1 << i
        } else if x % 3 == 2 {
            b |= 1 << i
        }
    }

    cnt1 := [N]int{}
    cnt2 := [N]int{}
    for _, x := range nums {
        if (x & int(a) == int(a)) && (x & int(b) == 0) {
            for i := 0; i < N; i++ {
                if x & (1 << i) != 0 {
                    cnt1[i]++
                }
            }
        } else if (x & int(b) == int(b)) && (x & int(a) == 0) {
            for i := 0; i < N; i++ {
                if x & (1 << i) != 0 {
                    cnt2[i]++
                }
            }
        }
    }

    for i := 0; i < N; i++ {
        if cnt1[i] % 3 == 1 {
            a |= 1 << i
        }
        if cnt2[i] % 3 == 2 {
            b |= 1 << i
        }
    }

    return []int{int(a), int(b)}
}