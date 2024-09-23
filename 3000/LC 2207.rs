impl Solution {
    pub fn maximum_subsequence_count(s: String, t: String) -> i64 {
        let mut ans = 0;
        let (mut cntA, mut cntB) = (0, 0);

        for c in s.chars() {
            if c == t.chars().nth(1).unwrap() {
                ans += cntA;
                cntB += 1;
            }

            if c == t.chars().nth(0).unwrap() {
                cntA += 1;
            }
        }

        ans + cntA.max(cntB)
    }
}