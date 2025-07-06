
import "strconv"
func closestFair(n int) int {
    m := map[int]int {
        1: 10,
        3: 1001,
        5: 100011,
        7: 10000111,
        9: 1000001111,
    }

    d := len(strconv.Itoa(n))
    if d & 1 == 1 {
        return m[d]
    }

    for !check(n) {
        n++
        d := len(strconv.Itoa(n))
        if d & 1 == 1 {
            return m[d]
        }
    }
    
    return n
}

func check(x int) bool {
    a, b := 0, 0
    for x > 0 {
        if x & 1 == 0 {
            a++
        } else {
            b++
        }
        x /= 10
    }

    return a == b
}