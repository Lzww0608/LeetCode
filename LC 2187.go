func minimumTime(time []int, totalTrips int) int64 {
    return int64(sort.Search(totalTrips * 1e7, func(x int) bool {
        res := 0
        for _, t := range time {
            res +=  x / t
            if res >= totalTrips {
                return true
            }
        }
        return false
    }))
}