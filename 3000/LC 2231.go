import "strconv"
func largestInteger(num int) int {
    s := []byte(strconv.Itoa(num))

    odd, even := []byte{}, []byte{}
    for _, c := range s {
        if int(c - '0') & 1 == 0 {
            even = append(even, c)
        } else {
            odd = append(odd, c)
        }
    }

    sort.Slice(odd, func(i, j int) bool {
        return odd[i] > odd[j]
    })
    sort.Slice(even, func(i, j int) bool {
        return even[i] > even[j]
    })

    for i, c := range s {
        if int(c - '0') & 1 == 0 {
            s[i] = even[0]
            even = even[1:]
        } else {
            s[i] = odd[0]
            odd = odd[1:]
        }
    }

    ans, _ := strconv.Atoi(string(s))
    return ans
}