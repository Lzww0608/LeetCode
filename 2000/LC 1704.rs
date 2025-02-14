use std::collections::HashMap;
impl Solution {
    pub fn halves_are_alike(s: String) -> bool {
        let s = s.as_bytes();
        let n = s.len();

        let mut cnt = HashMap::new();
        cnt.insert(b'a', true);
        cnt.insert(b'e', true);
        cnt.insert(b'i', true);
        cnt.insert(b'o', true);
        cnt.insert(b'u', true);
        cnt.insert(b'A', true);
        cnt.insert(b'E', true);
        cnt.insert(b'I', true);
        cnt.insert(b'O', true);
        cnt.insert(b'U', true);

        let mut a = 0;
        for i in 0..n {
            if cnt.contains_key(&s[i]) {
                if i < n / 2 {
                    a += 1;
                } else {
                    a -= 1;
                }
            }
        }

        a == 0
    }
}