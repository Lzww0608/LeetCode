func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    if s == goal {
        return checkEqual(s)
    }
    diff := [][2]int{}
    for i := range s {
        if s[i] != goal[i] {
            diff = append(diff, [2]int{int(s[i]-'a'), int(goal[i]-'a')})
        }
    }
    return len(diff) == 2 && diff[0][0] == diff[1][1] && diff[0][1] == diff[1][0]
}

func checkEqual(s string) bool {
    alp := [26]int{}
    for _, c := range s {
        alp[int(c-'a')]++
        if alp[int(c-'a')] > 1 {
            return true
        }
    }
    return false
}
