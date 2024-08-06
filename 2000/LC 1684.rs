impl Solution {
    pub fn count_consistent_strings(allowed: String, words: Vec<String>) -> i32 {
        let mut ans: i32 = 0;
        let mut mask: i32 = 0;

        for c in allowed.chars() {
            let x: i32 = (c as u8 - b'a') as i32;
            mask |= 1 << x;
        }

        for s in &words {
            let mut f: bool = true;
            for c in s.chars() {
                let x: i32 = (c as u8 - b'a') as i32;
                if mask & (1 << x) == 0 {
                    f = false;
                    break;
                }
            }
            if f {
                ans += 1;
            }
        }

        ans
    }
}