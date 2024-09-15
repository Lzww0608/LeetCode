impl Solution {
    pub fn count_letters(s: String) -> i32 {
        let mut ans = 0;
        let mut l = 0;
        for r in 0..s.len() {
            if s.chars().nth(r) != s.chars().nth(l) {
                l = r;
            }

            ans += (r - l + 1) as i32;
        }

        ans
    }
}