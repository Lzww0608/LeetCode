func checkRecord(s string) bool {
    return strings.Count(s, "A") <= 1 && !strings.Contains(s, "LLL")
}
