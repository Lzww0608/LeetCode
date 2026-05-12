func scoreValidator(events []string) []int {
    score, counter := 0, 0
    for _, v := range events {
        if v == "W" {
            counter++
        } else if v == "WD" || v == "NB" {
            score++
        } else {
            x, _ := strconv.Atoi(v)
            score += x
        }

        if counter >= 10 {
            break
        }
    }

    return []int{score, counter}
}