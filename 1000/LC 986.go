func intervalIntersection(firstList [][]int, secondList [][]int) (ans [][]int) {
    m, n := len(firstList), len(secondList)
    for i, j := 0, 0; i < m && j < n; {
        if firstList[i][0] <= secondList[j][0] {
            if firstList[i][1] < secondList[j][0] {
                i++
            } else {
                ans = append(ans, []int{secondList[j][0], min(firstList[i][1], secondList[j][1])})
                if firstList[i][1] <= secondList[j][1] {
                    i++
                } else {
                    j++
                }
            }
        } else {
            if secondList[j][1] < firstList[i][0] {
                j++
            } else {
                ans = append(ans, []int{firstList[i][0], min(firstList[i][1], secondList[j][1])})
                if secondList[j][1] <= firstList[i][1] {
                    j++
                } else {
                    i++
                }
            }
        }
    }

    return
}