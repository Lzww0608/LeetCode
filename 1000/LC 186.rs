impl Solution {
    pub fn reverse_words(s: &mut Vec<char>) {
        let n = s.len();

        let mut r = 0;
        while r < n {
            let l = r;
            while r < n && s[r] != ' ' {
                r += 1;
            }
            s[l..r].reverse();
            r += 1;
        }
        s.reverse();

    }
}