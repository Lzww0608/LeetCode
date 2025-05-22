const N int = 26
func slowestKey(releaseTimes []int, keysPressed string) byte {
    time := [N]int{}
    pre := 0
    for i := range keysPressed {
        x := int(keysPressed[i] - 'a')
        time[x] = max(time[x], releaseTimes[i] - pre)
        pre = releaseTimes[i]
    }

    id := 0 
    for i := range N {
        if time[i] >= time[id] {
            id = i
        }
    }

    return byte(id + 'a')
}