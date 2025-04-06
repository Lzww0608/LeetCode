func countMatches(items [][]string, ruleKey string, ruleValue string) (ans int) {
    for _, v := range items {
        if ruleKey == "color" && ruleValue == v[1] {
            ans++
        } else if ruleKey == "type" && ruleValue == v[0] {
            ans++
        } else if ruleKey == "name" && ruleValue == v[2] {
            ans++
        }
    }

    return
}