func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    if n == 0 {
        return 0
    }

    // Create a slice to store position and time to reach the target for each car
    cars := make([][2]float64, n)
    for i := 0; i < n; i++ {
        cars[i] = [2]float64{float64(position[i]), float64(target-position[i]) / float64(speed[i])}
    }

    // Sort cars based on position
    sort.Slice(cars, func(i, j int) bool {
        return cars[i][0] < cars[j][0]
    })

    fleets := 0
    currTime := 0.0
    // Iterate from the car closest to the target to the one farthest
    for i := n - 1; i >= 0; i-- {
        if cars[i][1] > currTime {
            fleets++
            currTime = cars[i][1]
        }
    }

    return fleets
}
