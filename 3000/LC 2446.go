func haveConflict(event1 []string, event2 []string) bool {
    l1 := solve(event1[0])
    r1 := solve(event1[1])
    l2 := solve(event2[0])
    r2 := solve(event2[1])
    
    return l1 <= l2 && r1 >= l2 || l2 <= l1 && r2 >= l1
}

func solve(str string) (res int) {
    s := strings.Split(str, ":")
    a, _ := strconv.Atoi(s[0])
    b, _ := strconv.Atoi(s[1])
    res = a * 60 + b 
    return
}