use std::collections::HashMap;

impl Solution {
    pub fn find_matrix(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut m: HashMap<i32, i32> = HashMap::new();
        for &x in &nums {
            *m.entry(x).or_insert(0) += 1;
        }

        let mut ans = Vec::new();
        while !m.is_empty() {
            let mut tmp = Vec::with_capacity(m.len());
            let mut keyToMove = Vec::new();

            for (&i, &mut x) in m.iter_mut() {
                tmp.push(i);
                if x - 1 == 0 {
                    keyToMove.push(i);
                } 
            }

            for x in keyToMove {
                m.remove(&x);
            }

            for key in tmp.iter() {
                if let Some(val) = m.get_mut(key) {
                    *val -= 1;
                }
            }
            ans.push(tmp);
        }

        ans
    }
}