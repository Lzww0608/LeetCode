use std::collections::HashMap;
impl Solution {
    pub fn array_change(nums: Vec<i32>, operations: Vec<Vec<i32>>) -> Vec<i32> {
        let mut m: HashMap<i32, i32> = HashMap::new();
        for op in operations.iter().rev() {
            let (x, mut y) = (op[0], op[1]);
            if let Some(&yy) = m.get(&y) {
                y = yy;
            }
            m.insert(x, y);
        }

        nums.into_iter().map(|num| *m.get(&num).unwrap_or(&num)).collect()
    }
}