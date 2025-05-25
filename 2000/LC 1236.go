/**
 * // This is HtmlParser's API interface.
 * // You should not implement it, or speculate about its implementation
 * type HtmlParser struct {
 *     func GetUrls(url string) []string {}
 * }
 */

func crawl(startUrl string, htmlParser HtmlParser) (ans []string) {
    q := []string{startUrl}
    m := make(map[string]bool)
    m[startUrl] = true

    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        ans = append(ans, cur)
        t := parse(cur)
        for _, s := range htmlParser.GetUrls(cur) {
            if !m[s] && parse(s) == t {
                m[s] = true
                q = append(q, s)
            }
        }
    }
    
    return ans
}

func parse(url string) string {
    if len(url) < 7 {
        return ""
    }

    url = url[7:]
    pos := strings.Index(url, "/")
    if pos == -1 {
        return url
    }

    return url[:pos]
}