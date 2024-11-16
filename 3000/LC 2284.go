
import "strings"
func largestWordCount(messages []string, senders []string) string {
    m := map[string]int{}
    ans := senders[0]
    for i, message := range messages {
        cnt := strings.Count(message, " ") + 1
        if m[senders[i]] += cnt; m[senders[i]] > m[ans] || (m[senders[i]] == m[ans] && ans < senders[i]) {
            ans = senders[i]
        }
    }

    return ans
}