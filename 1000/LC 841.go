func canVisitAllRooms(rooms [][]int) bool {
    n := len(rooms)
    col := make([]int, n)
    q := []int{}
    col[0] = 1
    for _, x := range rooms[0] {
        q = append(q, x)
    }
    for len(q) > 0 {
        t := q[0]
        q = q[1:]
        col[t] = 1
        for _, x := range rooms[t] {
            if col[x] == 0 {
                q = append(q, x)
            }
        }
    }
    for _, x := range col {
        if x == 0 {
            return false
        }
    }
    return true
}
