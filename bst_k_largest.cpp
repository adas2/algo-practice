/**** Find the k-th largest node in a BST ****/
#include <iostream>
#include <climits>

// #define INT_MIN -99999

using namespace std;

//<template class T>
struct t_node
{
  int val;
  struct t_node *left;
  struct t_node *right;
  // for order statistic tree
  // num nodes for the subtree rooted at this node
  int count;
};

// typedef struct t_node *tree;

class Solution
{
public:
  bool addNode(int val, struct t_node *&root);
  void initTree();
  // void lowest_common_ancestor(struct t_node root, int val1, int val2 );
  void printReverseInorder(struct t_node *root, int &count, int k, int &result);
  int find_k_largest_node(struct t_node *root, int k); // O(n) implementation
  int updateCount(struct t_node *root);
  int find_k_smallest(struct t_node *root, int k); // O(h) implementation
  // void printInorder(tree s);
  struct t_node *root;
};

void Solution::initTree()
{
  root = NULL;
}

bool Solution::addNode(int val, struct t_node *&root)
{
  if (!root)
  {
    // cout << "Creating new node" << endl;
    struct t_node *t = new (struct t_node);
    t->val = val;
    t->count = -1; // intiialize count
    t->left = t->right = NULL;
    root = t;

    return true;
  }
  // cout << "Root not NULL" << endl;
  if (val < root->val)
    return addNode(val, root->left);
  else
    return addNode(val, root->right);
}

// helper function
// print in the order when it prints highest first, count till k and then return node value
void Solution::printReverseInorder(struct t_node *root, int &cnt, int k, int &result)
{
  if (!root)
    return;
  // print right subtree
  printReverseInorder(root->right, cnt, k, result);
  // do operation with root
  cnt++;
  if (cnt == k)
  {
    result = root->val;
    return;
  }
  // cout << root->val;
  // cout << "  Count : " << cnt << endl;
  // print left value
  printReverseInorder(root->left, cnt, k, result);
}

// This is O(n) solution
int Solution::find_k_largest_node(struct t_node *root, int k)
{
  // e.g.if k=1 print rightmost element, k=3 third largest in reverse inorder sequence
  int count = 0;
  int result = INT_MIN; // some large value
  printReverseInorder(root, count, k, result);
  return result;
}

// NOTE: Use order statistics tree for O(h) solution
// calculate count-stats for each tree node
int Solution::updateCount(struct t_node *root)
{
  if (root == NULL)
  {
    return 0;
  }
  if (root->count != -1)
    return root->count;
  // else left count + right count + root
  root->count = updateCount(root->left) + updateCount(root->right) + 1;
  return root->count;
}

// find the k-th smallest node
int Solution::find_k_smallest(struct t_node *root, int k)
{
  if (root == NULL)
  {
    cout << "Root is NULL" << endl;
    return INT_MIN; // error case
  }
  // num nodes left subtree + root
  int lcount = root->count - ((!root->right) ? 0 : root->right->count);
  cout << "lcount " << lcount << endl;

  if (lcount == k)
    return root->val;
  if (lcount > k)
    return find_k_smallest(root->left, k);
  else // k > lcount
    return find_k_smallest(root->right, (k - lcount));
}

int main()
{
  Solution sol;
  sol.initTree();
  sol.addNode(6, sol.root);
  sol.addNode(3, sol.root);
  sol.addNode(9, sol.root);
  // sol.addNode(1, sol.root);
  sol.addNode(5, sol.root);
  sol.addNode(8, sol.root);
  sol.addNode(11, sol.root);
  sol.addNode(4, sol.root);
  sol.addNode(7, sol.root);
  sol.addNode(10, sol.root);
  // sol.addNode(12, sol.root);
  // cout << "In Order Traversal:" << endl;
  // sol.printInorder(sol.root);

  // sol.printReverseInorder(sol.root);
  int k = 3;
  cout << k << "-th Largest: " << sol.find_k_largest_node(sol.root, k) << endl;

  cout << sol.updateCount(sol.root) << endl;
  k = 5;
  cout << k << "-th Smallest: " << sol.find_k_smallest(sol.root, k) << endl;
  ;
}
