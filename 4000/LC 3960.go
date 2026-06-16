func getLength(nums []int) int {
    n := len(nums)
    mn, cur := 1, 1
    for i := 1; i < n; i++ {
        if nums[i] == nums[i - 1] {
            cur++
            mn = max(mn, cur)
        } else {
            cur = 1
        }
    } 

    for d := n; d > mn; d-- {
        cnt := make(map[int]int)
        freq := make(map[int]int)
        for i, x := range nums {
            if freq[cnt[x]] > 0 {
                if freq[cnt[x]]--; freq[cnt[x]] == 0 {
                    delete(freq, cnt[x])
                }
            }
            cnt[x]++
            freq[cnt[x]]++
            if i >= d - 1 {
                // if d == 8 && i == n - 2 {
                //     fmt.Println(cnt, freq)
                // }
                if len(freq) == 1 {
                    if len(cnt) == 1 {
                        return d
                    }
                } else if len(freq) == 2 {
                    tmp := []int{}
                    for k := range freq {
                        tmp = append(tmp, k)
                    }
                    if tmp[0] == tmp[1] * 2 || tmp[1] == tmp[0] * 2 {
                        return d
                    }
                }

                y := nums[i - d + 1]
                if freq[cnt[y]]--; freq[cnt[y]] == 0 {
                    delete(freq, cnt[y])
                }
                cnt[y]--
                if cnt[y] > 0 {
                    freq[cnt[y]]++
                }
            }
        }
    }

    return mn
}