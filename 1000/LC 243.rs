impl Solution {
    pub fn shortest_distance(words_dict: Vec<String>, word1: String, word2: String) -> i32 {
        let (mut a, mut b) = (i32::MIN / 2, i32::MIN / 2);
        let mut ans = i32::MAX;
        for (i, s) in words_dict.iter().enumerate() {
            if s == &word1 {
                ans = ans.min(i as i32 - b);
                a = i as i32;
            } else if s == &word2 {
                ans = ans.min(i as i32 - a);
                b = i as i32;
            }
        }

        ans
    }
}