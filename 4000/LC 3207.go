func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
    sort.Ints(enemyEnergies)
    if enemyEnergies[0] > currentEnergy {
        return 0
    }
    n := len(enemyEnergies)

    for i := n - 1; i >= 1; i-- {
        currentEnergy += enemyEnergies[i]
    }

    return int64(currentEnergy / enemyEnergies[0])
}