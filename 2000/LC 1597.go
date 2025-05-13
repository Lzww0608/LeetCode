/**
 * Definition for a binary tree node.
 * struct Node {
 *     char val;
 *     Node *left;
 *     Node *right;
 *     Node() : val(' '), left(nullptr), right(nullptr) {}
 *     Node(char x) : val(x), left(nullptr), right(nullptr) {}
 *     Node(char x, Node *left, Node *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
    int priority(char c) {
        if (c == '(') {
            return 4;
        } else if (c == '*' || c == '/') {
            return 3;
        } else if (c == '+' || c == '-') {
            return 2;
        } 
        return 1;
    }

public:
    Node* expTree(string s) {
        std::stack<Node*> nums;
        std::stack<char> ops;

        auto popOps = [&] () {
            auto right = nums.top();
            nums.pop();

            auto left = nums.top();
            nums.pop();
            auto root = new Node(ops.top(), left, right);
            ops.pop();
            nums.push(root);
        };

        for (char c : s) {
            if (c >= '0' && c <= '9') {
                nums.push(new Node(c));
            } else {
                if (ops.empty() || priority(ops.top()) < priority(c)) {
                    ops.push(c);
                } else {
                    while (!ops.empty() && ops.top() != '(' && priority(ops.top()) >= priority(c)) {
                        popOps();
                    }

                    if (c != ')') {
                        ops.push(c);
                    } else {
                        ops.pop();
                    }
                }
            }
        }

        while (!ops.empty()) {
            popOps();
        }

        return nums.top();
    }
};