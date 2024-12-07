func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) (ans int) {
    for i := min(startPos[0], homePos[0]); i <= max(startPos[0], homePos[0]); i++ {
        ans += rowCosts[i]
    }
    
    for j := min(startPos[1], homePos[1]); j <= max(startPos[1], homePos[1]); j++ {
        ans += colCosts[j]
    }

    return ans - rowCosts[startPos[0]] - colCosts[startPos[1]]
}