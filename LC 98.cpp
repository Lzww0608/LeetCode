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
    bool isValidBST(TreeNode* root) {
        long long t = (long long)INT_MIN - 1;
        function<bool(TreeNode*)> dfs = [&] (TreeNode* root) -> bool {
            if (root == nullptr) return true;
            if (!dfs(root->left)) return false;
            if (root->val <= t) return false;
            t = root->val; 
            return dfs(root->right);
        };
        return dfs(root);
    }
};