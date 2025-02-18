impl Solution {
    pub fn reorder_spaces(text: String) -> String {
        let cnt = text.chars().filter(|&c| c == ' ').count();
        let s: Vec<&str> = text.split_whitespace().collect();
        let n = s.len();

        if n <= 1 {
            return s.join("") + &" ".repeat(cnt);
        }

        s.join(&" ".repeat(cnt / (n - 1))) + &" ".repeat(cnt % (n - 1))
    }
}