impl Solution {
    pub fn find_words_containing(words: Vec<String>, x: char) -> Vec<i32> {
        let mut ans = Vec::new();
        for (i, s) in words.iter().enumerate() {
            if s.contains(x) {
                ans.push(i as i32);
            }
        }

        ans
    }
}