use std::collections::HashMap;
impl Solution {
    pub fn minimum_operations(nums: Vec<i32>) -> i32 {
        let n = nums.len() as i32;
        let mut m1: HashMap<i32, i32> = HashMap::new();
        let mut m2: HashMap<i32, i32> = HashMap::new();

        for (i, &x) in nums.iter().enumerate() {
            if (i & 1) == 0 {
                *m1.entry(x).or_insert(0) += 1;
            } else {
                *m2.entry(x).or_insert(0) += 1;
            }
        }

        let (mut a, mut b, mut c, mut d) = (-1, -2, -3, -4);
        for (&k, &v) in &m1 {
            if v > *m1.get(&a).unwrap_or(&0) {
                b = a;
                a = k;
            } else if v > *m1.get(&b).unwrap_or(&0) {
                b = k;
            }
        }
        for (&k, &v) in &m2 {
            if v > *m2.get(&c).unwrap_or(&0) {
                d = c;
                c = k;
            } else if v > *m2.get(&d).unwrap_or(&0) {
                d = k;
            }
        }

        if a != c {
            return n - *m1.get(&a).unwrap_or(&0) - *m2.get(&c).unwrap_or(&0);
        }

        return n - (*m1.get(&a).unwrap_or(&0) + *m2.get(&d).unwrap_or(&0)).max(*m1.get(&b).unwrap_or(&0) + *m2.get(&c).unwrap_or(&0));
    }
}