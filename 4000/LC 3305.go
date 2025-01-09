const N int = 26
func countOfSubstrings(s string, k int) (ans int) {
    m := map[byte]bool {
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }

    n := len(s)
    vow1, cst1 := 0, 0
    vow2, cst2 := 0, 0

    cnt1 := make([]int, N)
    cnt2 := make([]int, N)
    for l1, l2, r := 0, 0, 0; r < n; r++ {
        x := int(s[r] - 'a')
        if m[s[r]] {
            if cnt1[x]++; cnt1[x] == 1 {
                vow1++
            } 
            if cnt2[x]++; cnt2[x] == 1 {
                vow2++
            } 
        } else {
            cnt1[x]++
            cst1++

            cnt2[x]++
            cst2++
        }

        for cst1 >= k && vow1 == 5 {
            y := int(s[l1] - 'a')
            if m[s[l1]] {
                if cnt1[y]--; cnt1[y] == 0 {
                    vow1--
                }
            } else {
                cnt1[y]--
                cst1--
            }
            l1++
        }


        for cst2 > k && vow2 == 5 {
            y := int(s[l2] - 'a')
            if m[s[l2]] {
                if cnt2[y]--; cnt2[y] == 0 {
                    vow2--
                }
            } else {
                cnt2[y]--
                cst2--
            }
            l2++
        }

        ans += l1 - l2
    }

    return
}