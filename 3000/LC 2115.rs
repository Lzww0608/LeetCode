use::std::collections::{HashMap, VecDeque};
impl Solution {
    pub fn find_all_recipes(recipes: Vec<String>, ingredients: Vec<Vec<String>>, supplies: Vec<String>) -> Vec<String> {
        let mut g: HashMap<String, Vec<String>> = HashMap::new();
        let mut indegree: HashMap<String, i32> = HashMap::new();
        let mut ans: Vec<String> = Vec::new();
        let mut q: VecDeque<String> = VecDeque::new();

        for (i, r) in recipes.iter().enumerate() {
            indegree.insert(r.clone(), ingredients[i].len() as i32);
            for s in &ingredients[i] {
                g.entry(s.clone()).or_default().push(r.clone());
            }
        }

        for s in supplies {
            q.push_back(s);
        }

        while let Some(cur) = q.pop_front() {
            if let Some(neighbors) = g.get(&cur) {
                for neighbor in neighbors {
                    if let Some(count) = indegree.get_mut(neighbor) {
                        *count -= 1;
                        if *count == 0 {
                            q.push_back(neighbor.clone());
                            ans.push(neighbor.clone());
                        }
                    }
                }
            }
        }

        ans
    }
}