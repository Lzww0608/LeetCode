const N int = 10
func digitCount(num string) bool {
    cnt := [N]int{}
    for i := range num {
        cnt[i] = int(num[i] - '0')
    }

    for _, c := range num {
        x := int(c - '0')
        cnt[x]--
    }

    for _, x := range cnt {
        if x != 0 {
            return false
        }
    }

    return true
}