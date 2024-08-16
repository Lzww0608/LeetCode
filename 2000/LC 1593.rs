use std::collections::HashSet;
impl Solution {
    pub fn max_unique_split(s: String) -> i32 {
        let mut ans = 0;
        let mut m = HashSet::new();
        let n = s.len();

        fn dfs(s: &str, idx: usize, cnt: i32, n: usize, ans: &mut i32, m: &mut HashSet<String>) {
            if idx == n {
                *ans = (*ans).max(cnt);
                return;
            }

            for i in idx..n {
                let cur = &s[idx..i+1];
                if !m.contains(cur) {
                    m.insert(cur.to_string());
                    dfs(s, i + 1, cnt + 1, n, ans, m);
                    m.remove(cur);
                }
            }
        }
        dfs(&s, 0, 0, n, &mut ans, &mut m);
        ans
    }
}