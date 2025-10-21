func countTime(time string) int {
    a, b := 1, 1

    if time[0] == '?' {
        if time[1] == '?' {
            a = 24
        } else if time[1] >= '4' {
            a = 2
        } else {
            a = 3
        }
    } else {
        if time[1] == '?' {
            if time[0] == '2' {
                a = 4
            } else {
                a = 10
            }
        } 
    }

    if time[3] == '?' {
        if time[4] == '?' {
            b = 60
        } else {
            b = 6
        }
    } else {
        if time[4] == '?' {
            b = 10
        }
    }

    return a * b
}