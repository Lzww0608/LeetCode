impl Solution {
    pub fn find_permutation_difference(s: String, t: String) -> i32 {
        let mut pos = vec![0; 26];
        for (i, c) in s.chars().enumerate() {
            let x = c as u8 - b'a';
            pos[x as usize] = i;
        }

        let mut ans = 0;
        for (i, c) in t.chars().enumerate() {
            let x = c as u8 - b'a';
            ans += (i as i32 - pos[x as usize] as i32).abs();
        }

        ans as i32
    }
}