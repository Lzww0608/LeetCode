func countPairs(nums []int) (ans int) {
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if check(nums[i], nums[j]) {
                ans++
            } 
        }
    }

    return 
}

func check(x, y int) bool {
    s := strconv.Itoa(x)
    t := strconv.Itoa(y)
    if len(s) < len(t) {
        s = strings.Repeat("0", len(t) - len(s)) + s
    } else {
        t = strings.Repeat("0", len(s) - len(t)) + t
    }

    n := len(s)
    a := []int{}
    for i := 0; i < n; i++ {
        if s[i] != t[i] {
            a = append(a, i)
        }
    }

    if len(a) == 0 {
        return true 
    } else if len(a) == 2 && s[a[0]] == t[a[1]] && s[a[1]] == t[a[0]] {
        return true
    }

    return false
}