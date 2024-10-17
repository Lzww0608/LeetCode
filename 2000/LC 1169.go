use std::collections::HashSet;

#[derive(Debug, Clone)]
struct Transaction {
    name: String,
    time: i32,
    amount: i32,
    city: String,
    original: String,
}

impl Transaction {
    pub fn from_str(transaction_str: &str) -> Self {
        let parts: Vec<&str> = transaction_str.split(',').collect();
        Transaction {
            name: parts[0].to_string(),
            time: parts[1].parse::<i32>().unwrap(),
            amount: parts[2].parse::<i32>().unwrap(),
            city: parts[3].to_string(),
            original: transaction_str.to_string(),
        }
    }
}

impl Solution {
    pub fn invalid_transactions(transactions: Vec<String>) -> Vec<String> {
        let mut parsed_transactions: Vec<Transaction> = Vec::new();
        for transaction_str in &transactions {
            parsed_transactions.push(Transaction::from_str(transaction_str));
        }

        let mut ans: HashSet<usize> = HashSet::new();
        for i in 0..parsed_transactions.len() {
            let a = &parsed_transactions[i];

            if a.amount > 1000 {
                ans.insert(i);
            }

            for j in (i + 1)..parsed_transactions.len() {
                let b = &parsed_transactions[j];

                if b.name == a.name && (a.time - b.time).abs() <= 60 && a.city != b.city {
                    ans.insert(i);
                    ans.insert(j);
                }
            }
        }

        ans
            .into_iter()
            .map(|idx| parsed_transactions[idx].original.clone())
            .collect()
    }
}