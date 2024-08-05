impl Solution {
    pub fn max_depth(s: String) -> i32 {
        let mut ans = 0;
        let mut cur = 0;
        for c in s.chars() {
            if c == '(' {
                cur += 1;
                ans = ans.max(cur);
            } else if c == ')' {
                cur -= 1;
            }
        }

        ans
    }
}