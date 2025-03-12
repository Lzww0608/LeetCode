func highFive(items [][]int) (ans [][]int) {
    sort.Slice(items, func(i, j int) bool {
        if items[i][0] == items[j][0] {
            return items[i][1] > items[j][1] 
        }
        return items[i][0] < items[j][0]
    })
    cnt, avg := 0, 0
    for i := range items {
        if cnt >= 5 && i > 0 && items[i][0] == items[i-1][0] {
            continue
        } else if i > 0 && items[i][0] != items[i-1][0] {
            cnt = 0
        }
        avg += items[i][1]
        if cnt++; cnt == 5 {
            ans = append(ans, []int{items[i][0], avg / 5})
            avg = 0
        }
    }
    
    return
}