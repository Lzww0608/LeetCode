func sortFeatures(features []string, responses []string) []string {
    cnt := make(map[string]int)
    for _, response := range responses {
        vis := map[string]struct{}{}
        for _, s := range strings.Split(response, " ") {
            if _, ok := vis[s]; !ok {
                vis[s] = struct{}{}
                cnt[s]++
            }
        }
    }

    sort.SliceStable(features, func(i, j int) bool {
        return cnt[features[i]] > cnt[features[j]]
    })

    return features
}