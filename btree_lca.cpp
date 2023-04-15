/**** Find LCA for Generic Binary Tree (not BST) ***/

#include <iostream>
#include <vector>

using namespace std;
/****
//<template class T>
struct t_node{
  int val;
  struct t_node *left;
  struct t_node *right;
};
****/

class T_node
{
public:
  // constructor
  T_node(int val, T_node *left, T_node *right)
  {
    this->value = val;
    this->left = left;
    this->right = right;
  }
  // private:
  int value;
  T_node *left;
  T_node *right;
};

// Tree solution class
class Solution
{
public:
  Solution(T_node *root);
  // linear time lowest common ancestor
  T_node *find_LCA(T_node *root, int val1, int val2, bool &v1, bool &v2);
  void printInorder(T_node *s);
  // find is node exists in tree rooted at s; linear time O(n)
  bool isNodeInTree(T_node *s, int val);
  // O(nlon) time : non-optimal
  void lowest_common_ancestor(T_node *root, int val1, int val2);
  // private:
  T_node *root;
};

// constructor
Solution::Solution(T_node *root)
{
  this->root = root;
}

void Solution::printInorder(T_node *s)
{
  if (!s)
    return;
  printInorder(s->left);
  cout << " " << s->value << endl;
  printInorder(s->right);
}

// Optimal soln: time complexity O(n) , single pass
T_node *Solution::find_LCA(T_node *root, int val1, int val2, bool &v1, bool &v2)
{
  // empty tree
  if (!root)
    return NULL;

  // if root value is same as val1 and val 2
  if (root->value == val1)
  {
    v1 = true;
    return root;
  }
  if (root->value == val2)
  {
    v2 = true;
    return root;
  }

  // find LCA for left and right subtree
  T_node *left_lca = find_LCA(root->left, val1, val2, v1, v2);
  T_node *right_lca = find_LCA(root->right, val1, val2, v1, v2);

  // if leftLCa and rightLC are both non-NULL root is LCA
  if (left_lca && right_lca)
    return root;
  // if left_LCA is not-NULL and both v1 and v2 are true, left is LCA
  else if (left_lca)
    return left_lca;
  else // if(right_LCA)
    return right_lca;
}

// Time complexity: O(h.n) where h is height f tree i.e. O(nlogn)
void Solution::lowest_common_ancestor(T_node *root, int val1, int val2)
{
  if (!root)
    return;
  // if val1 or val2 matches root return LCA
  if (root->value == val1 || root->value == val2)
  {
    cout << "LCA: " << root->value << endl;
    return;
  }
  // if val1 is in left subtree and val2 in right subtree OR viceversa, root is LCA
  if ((isNodeInTree(root->left, val1) && isNodeInTree(root->right, val2)) ||
      (isNodeInTree(root->right, val1) && isNodeInTree(root->left, val2)))
  {
    cout << "LCA: " << root->value << endl;
    return;
  }
  // if val1 and val2 in left subtree
  if (isNodeInTree(root->left, val1) && isNodeInTree(root->left, val2))
  {
    lowest_common_ancestor(root->left, val1, val2);
  }
  else
  {
    lowest_common_ancestor(root->right, val1, val2);
  }
}

// find if node is in tree
bool Solution::isNodeInTree(T_node *root, int val)
{
  if (!root)
    return false;
  if (root->value == val)
    return true;

  return (isNodeInTree(root->left, val) || isNodeInTree(root->right, val));
}

int main()
{

  // create the sequence of nodes
  T_node *mRoot = new T_node(3, new T_node(5, NULL, NULL), new T_node(2, new T_node(10, new T_node(1, NULL, NULL), new T_node(6, NULL, NULL)), new T_node(17, NULL, NULL)));

  Solution mTree(mRoot);

  cout << "In Order Traversal:" << endl;
  mTree.printInorder(mTree.root);
  // optimal find LCA
  bool v1 = false, v2 = false;
  int n1, n2;
  cout << "enter nodes:" << endl;
  cin >> n1 >> n2;
  T_node *lca = mTree.find_LCA(mRoot, n1, n2, v1, v2);
  // check if both nodes are present in tree
  // NOTE: an extra pass is used here O(n)
  if ((v1 && v2) || (v1 && mTree.isNodeInTree(mRoot, n2)) || (v2 && mTree.isNodeInTree(mRoot, n1)))
    cout << "LCA: " << lca->value << endl;
  else
    cout << "LCA: NULL" << endl;

  // mTree.lowest_common_ancestor(mTree.root, 1, 17);
}
