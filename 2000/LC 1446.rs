

impl Solution {
    pub fn max_power(s: String) -> i32 {
        let mut ans: i32 = 1;
        let mut cur: i32 = 1;
        let mut pre = ' ';
        for c in s.chars() {
            if pre == c {
                cur += 1;
                ans = ans.max(cur);
            } else {
                pre = c;
                cur = 1;
            }
        }

        ans
    }
}