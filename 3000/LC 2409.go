var date [13]int = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    l := calc(max(arriveAlice, arriveBob))
    r := calc(min(leaveAlice, leaveBob))
    return max(r - l + 1, 0)
}

func calc(s string) (ans int) {
    month := int(s[0] - '0') * 10 + int(s[1] - '0')
    day := int(s[3] - '0') * 10 + int(s[4] - '0')

    for i := 1; i < month; i++ {
        ans += date[i]
    }
    
    return ans + day
}


