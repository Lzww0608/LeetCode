
func countSeniors(details []string) (ans int) {
    for _, detail := range details {
        n := len(detail)
        x, _ := strconv.Atoi(detail[n-4:n-2])
        if x > 60 {
            ans++
        }
    }

    return
}