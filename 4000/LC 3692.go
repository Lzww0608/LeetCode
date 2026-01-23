const N int = 26
func majorityFrequencyGroup(s string) string {
    cnt := [N]int{}
    for i := range s {
        x := int(s[i] - 'a')
        cnt[x]++
    }

    mx, k := 0, 0
    freq := make(map[int]int)
    for _, x := range cnt {
        if x == 0 {
            continue
        }
        freq[x]++

        if freq[x] >= mx {
            if freq[x] > mx {
                k = x
            } else {
                k = max(k, x)
            }
            mx = freq[x]
        }
    }

    ans := []byte{}
    for i, x := range cnt {
        if freq[x] == mx && k == x {
            ans = append(ans, byte('a' + i))
        }
    }

    return string(ans)
}