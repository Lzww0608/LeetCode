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
    TreeNode* subtreeWithAllDeepest(TreeNode* root) {
        TreeNode* ans = nullptr;
        int maxDepth = -1;
        function<int(TreeNode*, int)> dfs = [&] (TreeNode* root, int depth) -> int {
            if (root == nullptr) {
                maxDepth = max(maxDepth, depth);
                return depth;
            }
            int leftMaxDepth = dfs(root->left, depth + 1);
            int rightMaxDepth = dfs(root->right, depth + 1);
            if (leftMaxDepth == rightMaxDepth && leftMaxDepth == maxDepth) {
                ans = root;
            }
            return max(leftMaxDepth, rightMaxDepth);
        };
        dfs(root, 0);
        return ans;
    }
};