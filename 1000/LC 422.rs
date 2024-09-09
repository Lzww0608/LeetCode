impl Solution {
    pub fn valid_word_square(words: Vec<String>) -> bool {
        let n = words.len();

        for j in 0..n {
            let mut a = String::new();
            for i in 0..words.len() {
                if j < words[i].len() {
                    a.push(words[i].chars().nth(j).unwrap());
                } else {
                    break;
                }
            }
            if a != words[j] {
                return false;
            }
        }

        true
    }
}