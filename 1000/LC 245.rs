impl Solution {
    pub fn shortest_word_distance(words_dict: Vec<String>, s: String, t: String) -> i32 {
        let mut i = i32::MIN / 2;
        let mut j = i32::MIN / 2;
        let mut ans = i32::MAX;

        for (k, word) in words_dict.iter().enumerate() {
            if *word == s {
                i = k as i32;
                ans = ans.min(i - j);
                if s == t {
                    std::mem::swap(&mut i, &mut j);
                }
            } else if *word == t {
                j = k as i32;
                ans = ans.min(j - i);
            }
        } 

        ans
    }
}