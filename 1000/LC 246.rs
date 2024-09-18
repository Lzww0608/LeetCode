use std::collections::HashMap;
impl Solution {
    pub fn is_strobogrammatic(num: String) -> bool {
        let m = vec![
            ('9', '6'),
            ('6', '9'),
            ('1', '1'),
            ('0', '0'),
            ('8', '8'),
        ].into_iter().collect::<HashMap<char, char>>();

        let n = num.len();

        for i in 0..=n/2 {
            if m.get(&num.chars().nth(i).unwrap()) != Some(&num.chars().nth(n-1-i).unwrap()) {
                return false;
            }
        }

        true
    }
}