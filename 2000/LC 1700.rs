impl Solution {
    pub fn count_students(students: Vec<i32>, sandwiches: Vec<i32>) -> i32 {
        let mut cirle: i32 = 0;
        let mut square: i32 = 0;
        for &x in &students {
            if x == 1 {
                cirle += 1;
            } else {
                square += 1;
            }
        }

        for (i, &x) in sandwiches.iter().enumerate() {
            if x == 1 {
                cirle -= 1;
            } else {
                square -= 1;
            }

            if cirle < 0 || square < 0 {
                return (sandwiches.len() - i) as i32
            }
        }

        0
    }
}