func countMentions(numberOfUsers int, events [][]string) []int {
    sort.Slice(events, func(i, j int) bool {
        x, _ := strconv.Atoi(events[i][1])
        y, _ := strconv.Atoi(events[j][1])
        if x == y {
            return events[i][0][0] != 'M'
        }
        return x < y
    })

    ans := make([]int, numberOfUsers)
    online := make([]int, numberOfUsers)
    for _, v := range events {
        time, _ := strconv.Atoi(v[1])
        if (v[0][0] == 'O') {
            x, _ := strconv.Atoi(v[2])
            online[x] = time + 60
        } else if (v[2][0] == 'A') {
            for i := range ans {
               ans[i]++ 
            }
        } else if (v[2][0] == 'H') {
            for i, x := range online {
                if x <= time {
                    ans[i]++
                }
            }
        } else {
            s := strings.Split(v[2], " ")
            for _, t := range s {
                x, _ := strconv.Atoi(t[2:])
                ans[x]++
            }
        }
    }

    return ans
}