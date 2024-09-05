use std::collections::HashMap;
impl Solution {
    pub fn length_of_longest_substring_two_distinct(s: String) -> i32 {
        let mut cnt = HashMap::new();
        let mut sum = 0;
        let mut ans = 0;
        let n = s.len();
        let s = s.as_bytes();
        let (mut l, mut r) = (0, 0);

        while r < n {
            let x = s[r];
            *cnt.entry(x).or_insert(0) += 1;

            if cnt[&x] == 1 {
                sum += 1;
            }

            while sum > 2 {
                if let Some(t) = cnt.get_mut(&s[l]) {
                    *t -= 1;
                    if *t == 0 {
                        sum -= 1;
                    }
                }
                l += 1;
            }

            ans = ans.max(r - l + 1);
            r += 1;
        }

        ans as i32
    }
}