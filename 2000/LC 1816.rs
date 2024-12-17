impl Solution {
    pub fn truncate_sentence(s: String, k: i32) -> String {
        let words: Vec<&str> = s.split_whitespace().collect();
        words.into_iter().take(k as usize).collect::<Vec<&str>>().join(" ")
    }
}