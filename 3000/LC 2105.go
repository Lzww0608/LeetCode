func minimumRefill(plants []int, capacityA int, capacityB int) (ans int) {
    n := len(plants)
    a, b := capacityA, capacityB
    for i, j := 0, n - 1; i <= j; i, j = i + 1, j - 1 {
        if i == j {
            if a < plants[i] && b < plants[i] {
                ans++
            }
        } else {
            if a < plants[i] {
                a = capacityA - plants[i]
                ans++
            } else {
                a -= plants[i]
            }

            if b < plants[j] {
                b = capacityB - plants[j]
                ans++
            } else {
                b -= plants[j]
            }
        }
    }

    return 
}
