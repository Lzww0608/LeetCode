/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* replaceValueInTree(TreeNode* root) {
        root->val = 0;
        vector<TreeNode*> q{root};
        while (!q.empty()) {
            vector<TreeNode*> next;
            int sum = 0;
            for (auto& node : q) {
                if (node->left != nullptr)  {
                    next.push_back(node->left);
                    sum += node->left->val;
                }
                if (node->right != nullptr) {
                    next.push_back(node->right); 
                    sum += node->right->val;
                }
            }
            for (auto& node : q) {
                int sum_child = (node->left == nullptr ? 0 : node->left->val) + (node->right == nullptr ? 0 : node->right->val);
                if (node->left != nullptr) node->left->val = sum - sum_child;
                if (node->right != nullptr) node->right->val = sum - sum_child;
            }
            q = move(next);
        }
        return root;
    }
};