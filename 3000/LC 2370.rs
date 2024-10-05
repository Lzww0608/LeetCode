use std::cmp::{min, max};
impl Solution {
    pub fn longest_ideal_string(mut s: String, k: i32) -> i32 {
        let mut f = vec![0; 26];
        let s = s.as_bytes();
        for i in 0..s.len() {
            let x = (s[i] & 31) as i32 - 1;
            let l = max(0, x - k) as usize;
            let r = min(25, x + k) as usize;
            let mut mx = 0;
            for j in l..=r {
                mx = max(mx, f[j] + 1);
            }
            f[x as usize] = mx;
            //println!("{}, {}, {}, {}, {}", i, x, l, r, mx);
        }

        *f.iter().max().unwrap()
    }
}