func splitMessage(message string, limit int) []string {
    for i, cap, tail := 1, 0, 0; ; i++ {
        if i < 10 {
            tail = 5
        } else if i < 100 {
            if i == 10 {
                cap -= 9
            }
            tail = 7
        } else if i < 1000 {
            if i == 100 {
                cap -= 99
            }
            tail = 9
        } else {
            if i == 1000 {
                cap -= 999
            }
            tail = 11
        }
        if tail >= limit {
            return nil
        }
        cap += limit - tail
        if cap < len(message) {
            continue
        }

        ans := make([]string, 0, i)
        for j := 1; j <= i; j++ {
            s := fmt.Sprintf("<%d/%d>", j, i)
            m := min(len(message), limit - len(s))
            ans = append(ans, message[:m] + s)
            message = message[m:]
        }
        return ans
    }

    return nil
}