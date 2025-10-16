const N int = 26
func equalFrequency(word string) bool {
    cnt := [N]int{}
    for _, c := range word {
        cnt[int(c - 'a')]++
    }

    a := []int{}
    for _, x := range cnt {
        if x != 0 {
            a = append(a, x)
        }
    }
    sort.Ints(a)
    n := len(a)
    return n == 1 || a[0] == 1 && check(a[1:]) || a[n - 1] - 1 == a[n - 2] && check(a[:n - 1])
}

func check(a []int) bool {
    n := len(a)
    for i := 1; i < n; i++ {
        if a[i] != a[0] {
            return false
        }
    }

    return true
}