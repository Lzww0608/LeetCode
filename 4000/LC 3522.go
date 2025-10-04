func calculateScore(instructions []string, values []int) int64 {
    ans := 0
    n := len(instructions)
    for i := 0; i >= 0 && i < n && instructions[i] != ""; {
        if instructions[i] == "add" {
            ans += values[i]
            instructions[i] = ""
            i++
        } else {
            instructions[i] = ""
            i += values[i]
        }
    }

    return int64(ans)
}