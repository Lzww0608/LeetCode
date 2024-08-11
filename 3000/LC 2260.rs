use::std::collections::HashMap;
impl Solution {
    pub fn minimum_card_pickup(cards: Vec<i32>) -> i32 {
        let mut ans: i32 = std::i32::MAX;
        let mut m = HashMap::new();
        for (i, &x) in cards.iter().enumerate() {
            if let Some(&v) = m.get(&x) {
                ans = ans.min((i - v + 1) as i32);
            }
            m.insert(x, i);
        }

        if ans == std::i32::MAX {
            return -1;
        }

        ans
    }
}