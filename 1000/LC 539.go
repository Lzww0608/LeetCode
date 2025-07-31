
import "math"
func findMinDifference(timePoints []string) int {
    sort.Strings(timePoints)
    ans := math.MaxInt32 
    n := len(timePoints)
    for i := 0; i < n; i++ {
        ans = min(ans, solve(timePoints[i], timePoints[(i - 1 + n) % n]))
    }

import "math"
func findMinDifference(timePoints []string) int {
    if len(timePoints) > 24 * 60 {
        return 0
    }
    sort.Strings(timePoints)
    ans := math.MaxInt32 
    n := len(timePoints)
    for i := 0; i < n; i++ {
        ans = min(ans, solve(timePoints[i], timePoints[(i - 1 + n) % n]))
    }

    return ans
}

func solve(s, t string) int {
    a := strings.Split(s, ":")
    b := strings.Split(t, ":")
    x1, _ := strconv.Atoi(a[0])
    y1, _ := strconv.Atoi(a[1])
    x2, _ := strconv.Atoi(b[0])
    y2, _ := strconv.Atoi(b[1])

    x := x1 * 60 + y1 
    y := x2 * 60 + y2 
    return min(abs(x - y), abs(x + 24 * 60 - y))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}
    return ans
}

func solve(s, t string) int {
    a := strings.Split(s, ":")
    b := strings.Split(t, ":")
    x1, _ := strconv.Atoi(a[0])
    y1, _ := strconv.Atoi(a[1])
    x2, _ := strconv.Atoi(b[0])
    y2, _ := strconv.Atoi(b[1])

    x := x1 * 60 + y1 
    y := x2 * 60 + y2 
    return min(abs(x - y), abs(x + 24 * 60 - y))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}