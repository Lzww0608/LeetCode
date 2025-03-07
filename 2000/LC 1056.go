func confusingNumber(n int) bool {
    rotate := map[int]int {
        0: 0,
        1: 1,
        6: 9,
        8: 8,
        9: 6,
    }

    m := 0
    for t := n; t > 0; t /= 10 {
        y := t % 10
        if x, ok := rotate[y]; !ok {
            return false
        }  else {
            m = m * 10 + x 
        }
    } 

    return m != n
}