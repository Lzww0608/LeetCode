impl Solution {
    pub fn clear_digits(s: String) -> String {
        let s = s.as_str();
        let mut st = Vec::new();

        for c in s.chars() {
            if c.is_digit(10) {
                if !st.is_empty() {
                    st.pop();
                }
            } else {
                st.push(c);
            }
        }
        st.into_iter().collect()
    }
}