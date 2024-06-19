func angleClock(hour int, minutes int) float64 {
    var a, b float64
    a = (float64(hour % 12) + float64(minutes) / 60) * 30
    b = float64(minutes) * 6

    if b > a {
        return min(b - a, a - b + 360)
    }

    return min(a - b, b - a + 360)
}