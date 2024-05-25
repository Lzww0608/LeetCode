func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    if xCenter >= x1 && xCenter <= x2 && yCenter >= y1 && yCenter <= y2 {
        return true
    }
    
    nearestX := math.Max(float64(x1), math.Min(float64(xCenter), float64(x2)))
    nearestY := math.Max(float64(y1), math.Min(float64(yCenter), float64(y2)))

    // 计算圆心到最近点的距离
    deltaX := float64(xCenter) - nearestX
    deltaY := float64(yCenter) - nearestY
    distanceSquared := deltaX * deltaX + deltaY * deltaY

    // 判断距离是否小于等于圆的半径
    return distanceSquared <= float64(radius * radius)
}