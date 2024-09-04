impl Solution {
    pub fn min_steps(s: String, t: String) -> i32 {
        let mut ans = 0;
        let mut cnt = vec![0i32; 26];

        for c in s.chars() {
            let mut x = c as u8 - b'a';
            cnt[x as usize] += 1;
        }

        for c in t.chars() {
            let mut x = c as u8 - b'a';
            cnt[x as usize] -= 1;
        }

        for &x in &cnt {
            ans += x.abs();
        }

        ans
    }
}